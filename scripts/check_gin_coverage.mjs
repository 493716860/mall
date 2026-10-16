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

// Extract frontend API paths
const feFiles = [...walk('apps/admin/src/api', ['.ts']), ...walk('apps/miniapp/src', ['.ts', '.vue', '.js'])];
const feApis = new Set();

for (const file of feFiles) {
  const content = fs.readFileSync(file, 'utf8');
  // url: '/admin-api/...' or url: '/app-api/...'
  const matches = content.matchAll(/url:\s*['"`]([^'"`]+)['"`]/g);
  for (const m of matches) {
    let raw = m[1].trim().split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
    if (raw.startsWith('/')) {
      feApis.add(raw);
    }
  }
}

// Extract Go routes
const goFiles = walk('apps/server-go', ['.go']);
const goRoutes = [];

for (const file of goFiles) {
  const content = fs.readFileSync(file, 'utf8');
  
  // Find group bases like r.Group("/admin-api/system/auth")
  const groupMatches = content.matchAll(/(\w+)\s*:=\s*r\.Group\("([^"]+)"\)/g);
  const groups = {};
  for (const gm of groupMatches) {
    groups[gm[1]] = gm[2];
  }

  // Find routes
  const routeMatches = content.matchAll(/(\w+)\.(GET|POST|PUT|DELETE)\("([^"]+)"/g);
  for (const rm of routeMatches) {
    const groupVar = rm[1];
    const method = rm[2];
    const subPath = rm[3];
    const basePath = groups[groupVar] || '';
    let full = (basePath + (subPath ? (subPath.startsWith('/') ? subPath : '/' + subPath) : '')).replace(/\/+/g, '/').replace(/\/+$/, '');
    goRoutes.push({ method, path: full });
  }
}

const goPathSet = new Set(goRoutes.map(r => r.path));

const matched = [];
const missing = [];

for (const url of feApis) {
  if (goPathSet.has(url)) {
    matched.push(url);
  } else {
    // Check if path without prefix matches
    const unPrefixed = url.replace(/^\/admin-api/, '').replace(/^\/app-api/, '');
    const hasUnprefixed = goRoutes.some(r => r.path.endsWith(unPrefixed));
    if (hasUnprefixed) {
      matched.push(url);
    } else {
      missing.push(url);
    }
  }
}

console.log('=== COVERAGE CHECK ===');
console.log('Total Frontend API Endpoints declared:', feApis.size);
console.log('Total Go Gin Routes registered:', goRoutes.length);
console.log('Directly Matched Endpoints:', matched.length);
console.log('Missing/Unimplemented Endpoints:', missing.length);

if (missing.length > 0) {
  console.log('\nSample Missing Endpoints:');
  missing.slice(0, 30).forEach(m => console.log('- ' + m));
}
