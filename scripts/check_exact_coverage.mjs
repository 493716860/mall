import fs from 'node:fs';
import path from 'node:path';

function walk(dir, ext) {
  let results = [];
  if (!fs.existsSync(dir)) return results;
  for (const f of fs.readdirSync(dir)) {
    const full = path.join(dir, f);
    const stat = fs.statSync(full);
    if (stat.isDirectory()) {
      results = results.concat(walk(full, ext));
    } else if (ext.some(e => f.endsWith(e))) {
      results.push(full);
    }
  }
  return results;
}

// 1. Admin endpoints -> prefixed with /admin-api
const adminFiles = walk('apps/admin/src/api', ['.ts']);
const adminApis = new Map();

for (const file of adminFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const matches = content.matchAll(/url:\s*[`'"]([^`'"]+)[`'"]/g);
  for (const m of matches) {
    let clean = m[1].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (clean.startsWith('/')) {
      const full = clean.startsWith('/admin-api') ? clean : '/admin-api' + clean;
      adminApis.set(full, file);
    }
  }
}

// 2. Miniapp endpoints -> prefixed with /app-api
const miniappFiles = walk('apps/miniapp/sheep/api', ['.js', '.ts']);
const miniappApis = new Map();

for (const file of miniappFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const matches = content.matchAll(/url:\s*[`'"]([^`'"]+)[`'"]/g);
  for (const m of matches) {
    let clean = m[1].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (clean.startsWith('/')) {
      const full = clean.startsWith('/app-api') ? clean : '/app-api' + clean;
      miniappApis.set(full, file);
    }
  }
}

// 3. Go Gin routes
const goFiles = walk('apps/server-go', ['.go']);
const goRoutes = new Set();

for (const file of goFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const groupMatches = content.matchAll(/(\w+)\s*:=\s*r\.Group\("([^"]+)"\)/g);
  const groups = {};
  for (const gm of groupMatches) groups[gm[1]] = gm[2];

  const routeMatches = content.matchAll(/(\w+)\.(GET|POST|PUT|DELETE)\("([^"]+)"/g);
  for (const rm of routeMatches) {
    const groupVar = rm[1];
    const subPath = rm[3];
    const basePath = groups[groupVar] || '';
    let full = (basePath + (subPath ? (subPath.startsWith('/') ? subPath : '/' + subPath) : '')).replace(/\/+/g, '/').replace(/\/+$/, '');
    goRoutes.add(full);
  }
}

const adminMissing = [];
for (const [api, file] of adminApis.entries()) {
  const isMatch = goRoutes.has(api) || [...goRoutes].some(r => r.replace(/\{[^}]+\}/g, '{id}') === api);
  if (!isMatch) adminMissing.push({ api, file });
}

const miniappMissing = [];
for (const [api, file] of miniappApis.entries()) {
  const isMatch = goRoutes.has(api) || [...goRoutes].some(r => r.replace(/\{[^}]+\}/g, '{id}') === api);
  if (!isMatch) miniappMissing.push({ api, file });
}

console.log('=== 精准覆盖率分析报告 ===');
console.log('【管理后台端 apps/admin】');
console.log(`- 前端调用接口总数: ${adminApis.size} 个`);
console.log(`- Gin 已实现接口:   ${adminApis.size - adminMissing.length} 个`);
console.log(`- 尚未匹配接口:     ${adminMissing.length} 个`);

console.log('\n【微信小程序端 apps/miniapp】');
console.log(`- 前端调用接口总数: ${miniappApis.size} 个`);
console.log(`- Gin 已实现接口:   ${miniappApis.size - miniappMissing.length} 个`);
console.log(`- 尚未匹配接口:     ${miniappMissing.length} 个`);

fs.writeFileSync('/tmp/exact_missing.json', JSON.stringify({
  adminMissing,
  miniappMissing
}, null, 2));

console.log('\n详情已写入 /tmp/exact_missing.json');
