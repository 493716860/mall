import fs from 'node:fs';
import path from 'node:path';

function walk(dir, ext) {
  let results = [];
  if (!fs.existsSync(dir)) return results;
  for (const f of fs.readdirSync(dir)) {
    const full = path.join(dir, f);
    const stat = fs.statSync(full);
    if (stat.isDirectory()) results = results.concat(walk(full, ext));
    else if (ext.some(e => f.endsWith(e))) results.push(full);
  }
  return results;
}

// 1. Admin endpoints
const adminFiles = walk('apps/admin/src/api', ['.ts']);
const adminCalls = new Map();

for (const file of adminFiles) {
  const content = fs.readFileSync(file, 'utf8');
  // Match request.get({ url: '...' }) etc.
  const regex = /request\.(get|post|put|delete|download)[\s\S]*?url:\s*[`'"]([^`'"]+)[`'"]/g;
  let m;
  while ((m = regex.exec(content)) !== null) {
    let method = m[1].toUpperCase();
    if (method === 'DOWNLOAD') method = 'GET';
    let url = m[2].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (!url.startsWith('/admin-api')) {
      url = '/admin-api' + (url.startsWith('/') ? url : '/' + url);
    }
    adminCalls.set(`${method} ${url}`, { method, url, file });
  }
}

// 2. Miniapp endpoints
const miniappFiles = walk('apps/miniapp/sheep/api', ['.js', '.ts']);
const miniappCalls = new Map();

for (const file of miniappFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const regex = /request\.(get|post|put|delete|download)[\s\S]*?url:\s*[`'"]([^`'"]+)[`'"]/g;
  let m;
  while ((m = regex.exec(content)) !== null) {
    let method = m[1].toUpperCase();
    if (method === 'DOWNLOAD') method = 'GET';
    let url = m[2].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (!url.startsWith('/app-api')) {
      url = '/app-api' + (url.startsWith('/') ? url : '/' + url);
    }
    miniappCalls.set(`${method} ${url}`, { method, url, file });
  }
  const objRegex = /url:\s*[`'"]([^`'"]+)[`'"][\s\S]*?method:\s*[`'"](GET|POST|PUT|DELETE)[`'"]/gi;
  while ((m = objRegex.exec(content)) !== null) {
    let method = m[2].toUpperCase();
    let url = m[1].split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (!url.startsWith('/app-api')) {
      url = '/app-api' + (url.startsWith('/') ? url : '/' + url);
    }
    miniappCalls.set(`${method} ${url}`, { method, url, file });
  }
}

// 3. Scan Go Gin routes
const goFiles = walk('apps/server-go', ['.go']);
const goRoutes = new Map();

for (const file of goFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const groupMatches = content.matchAll(/(\w+)\s*:=\s*r\.Group\("([^"]+)"\)/g);
  const groups = {};
  for (const gm of groupMatches) groups[gm[1]] = gm[2];

  const routeMatches = content.matchAll(/(\w+)\.(GET|POST|PUT|DELETE)\("([^"]+)"/g);
  for (const rm of routeMatches) {
    const groupVar = rm[1];
    const method = rm[2].toUpperCase();
    const subPath = rm[3];
    const basePath = groups[groupVar] || '';
    let full = (basePath + (subPath ? (subPath.startsWith('/') ? subPath : '/' + subPath) : '')).replace(/\/+/g, '/').replace(/\/+$/, '');
    goRoutes.set(`${method} ${full.replace(/:id/g, '{id}')}`, { method, path: full, file });
  }
}

console.log(`Frontend Admin endpoints: ${adminCalls.size}`);
console.log(`Frontend Miniapp endpoints: ${miniappCalls.size}`);
console.log(`Go Gin registered endpoints: ${goRoutes.size}`);

const adminMethodMismatches = [];
for (const [key, call] of adminCalls.entries()) {
  if (!goRoutes.has(key)) {
    // Check if the URL exists under a DIFFERENT method
    const matchedUrlRoutes = [...goRoutes.entries()].filter(([k, v]) => k.split(' ')[1] === call.url);
    adminMethodMismatches.push({
      required: key,
      url: call.url,
      method: call.method,
      file: call.file,
      existingInGo: matchedUrlRoutes.map(([k]) => k)
    });
  }
}

const miniappMethodMismatches = [];
for (const [key, call] of miniappCalls.entries()) {
  if (!goRoutes.has(key)) {
    const matchedUrlRoutes = [...goRoutes.entries()].filter(([k, v]) => k.split(' ')[1] === call.url);
    miniappMethodMismatches.push({
      required: key,
      url: call.url,
      method: call.method,
      file: call.file,
      existingInGo: matchedUrlRoutes.map(([k]) => k)
    });
  }
}

console.log(`Admin mismatches (method or route missing): ${adminMethodMismatches.length}`);
console.log(`Miniapp mismatches (method or route missing): ${miniappMethodMismatches.length}`);

fs.writeFileSync('/tmp/mismatches.json', JSON.stringify({
  adminMethodMismatches,
  miniappMethodMismatches
}, null, 2));

console.log('Results written to /tmp/mismatches.json');
