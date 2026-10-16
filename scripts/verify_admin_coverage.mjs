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

// 1. Scan apps/admin/src
const adminFiles = walk('apps/admin/src', ['.ts', '.js', '.vue']);
console.log('Total admin frontend source files scanned:', adminFiles.length);

const frontendApis = [];
for (const file of adminFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const relFile = file.replace('apps/admin/', '');
  
  // url: '...' or url: `...`
  const matches = content.matchAll(/url:\s*([`'"][^`'"]+[`'"])/g);
  for (const m of matches) {
    let raw = m[1].slice(1, -1).trim();
    if (raw.startsWith('/')) {
      let clean = raw.split('?')[0].replace(/\$\{[^}]+\}/g, '{id}').replace(/\/+$/, '');
      clean = clean.replace(/^\/admin-api/, '');
      frontendApis.push({ file: relFile, url: clean, raw });
    }
  }
}

const uniqueFrontendUrls = [...new Set(frontendApis.map(a => a.url))].filter(Boolean).sort();
console.log('Unique admin frontend API URLs found:', uniqueFrontendUrls.length);

// 2. Scan apps/server-java Admin controllers
const javaFiles = walk('apps/server-java', ['.java']).filter(f => f.includes('/controller/admin/'));
console.log('Total Admin controller Java files in server-java:', javaFiles.length);

const javaAdminEndpoints = [];
for (const file of javaFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const classMatch = content.match(/@RequestMapping\("([^"]+)"\)/);
  const basePath = classMatch ? classMatch[1] : '';
  
  const methodMatches = content.matchAll(/@(GetMapping|PostMapping|PutMapping|DeleteMapping|RequestMapping)\((?:value\s*=\s*)?(?:"([^"]*)")?/g);
  for (const m of methodMatches) {
    const httpMethod = m[1].replace('Mapping', '').toUpperCase();
    const subPath = m[2] !== undefined ? m[2] : '';
    let full = (basePath + (subPath ? (subPath.startsWith('/') ? subPath : '/' + subPath) : '')).replace(/\/+/g, '/').replace(/\/+$/, '');
    full = full.replace(/^\/admin-api/, '');
    const normalized = full.replace(/\{[^}]+\}/g, '{id}');
    javaAdminEndpoints.push({ file: path.basename(file), httpMethod, rawPath: full, normalized });
  }
}

console.log('Total Java Admin REST mappings found in server-java:', javaAdminEndpoints.length);

// 3. Match
const matched = [];
const missing = [];

for (const feUrl of uniqueFrontendUrls) {
  const isMatch = javaAdminEndpoints.some(j => {
    if (j.rawPath === feUrl || j.normalized === feUrl) return true;
    const jParts = j.normalized.split('/');
    const fParts = feUrl.split('/');
    if (jParts.length !== fParts.length) return false;
    return jParts.every((p, idx) => p === fParts[idx] || p === '{id}' || fParts[idx] === '{id}');
  });

  if (isMatch) {
    matched.push(feUrl);
  } else {
    // Find files where it was called
    const callers = frontendApis.filter(a => a.url === feUrl).map(a => a.file);
    missing.push({ url: feUrl, callers: [...new Set(callers)] });
  }
}

console.log('=== MATCH SUMMARY ===');
console.log('Fully Matched:', matched.length);
console.log('Missing / Differences:', missing.length);

if (missing.length > 0) {
  console.log('\nList of Missing/Different Endpoints:');
  missing.forEach(m => console.log(`- ${m.url} (called in: ${m.callers.join(', ')})`));
}

fs.writeFileSync('admin_coverage_report.json', JSON.stringify({ matched, missing, totalFrontend: uniqueFrontendUrls.length, totalJava: javaAdminEndpoints.length }, null, 2));
