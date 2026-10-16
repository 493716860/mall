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

// 1. Collect all admin API endpoints from apps/admin/src/api
const adminFiles = walk('apps/admin/src/api', ['.ts']);
const adminApis = new Map(); // path -> file

for (const file of adminFiles) {
  const content = fs.readFileSync(file, 'utf8');
  // Match url: '...' or url: `...` or url: "..."
  const matches = content.matchAll(/url:\s*[`'"]([^`'"]+)[`'"]/g);
  for (const m of matches) {
    let clean = m[1].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (clean.startsWith('/')) {
      adminApis.set(clean, file);
    }
  }
}

// 2. Collect all miniapp endpoints from apps/miniapp/src
const miniappFiles = walk('apps/miniapp/src', ['.ts', '.js', '.vue']);
const miniappApis = new Map();

for (const file of miniappFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const matches = content.matchAll(/url:\s*[`'"]([^`'"]+)[`'"]/g);
  for (const m of matches) {
    let clean = m[1].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (clean.startsWith('/app-api')) {
      miniappApis.set(clean, file);
    }
  }
}

console.log('Total Admin API endpoints in apps/admin/src/api:', adminApis.size);
console.log('Total Miniapp API endpoints in apps/miniapp/src:', miniappApis.size);

// Save to JSON for exact verification
const adminList = Array.from(adminApis.keys()).sort();
const miniappList = Array.from(miniappApis.keys()).sort();

fs.writeFileSync('/tmp/all_frontend_apis.json', JSON.stringify({
  adminCount: adminList.length,
  admin: adminList,
  miniappCount: miniappList.length,
  miniapp: miniappList
}, null, 2));

console.log('Saved /tmp/all_frontend_apis.json');
