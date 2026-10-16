import fs from 'node:fs';
import path from 'node:path';

function walk(dir) {
  let results = [];
  if (!fs.existsSync(dir)) return results;
  const list = fs.readdirSync(dir);
  for (const file of list) {
    const full = path.join(dir, file);
    const stat = fs.statSync(full);
    if (stat && stat.isDirectory()) {
      results = results.concat(walk(full));
    } else if (file.endsWith('.ts') || file.endsWith('.js') || file.endsWith('.vue')) {
      results.push(full);
    }
  }
  return results;
}

const adminFiles = walk('apps/admin/src');
const miniappFiles = walk('apps/miniapp');

const adminEndpoints = new Map();
const miniappEndpoints = new Map();

// Parse admin API files and Vue components
for (const file of adminFiles) {
  const content = fs.readFileSync(file, 'utf8');
  // Find all request calls
  const regex = /request\.(get|post|put|delete|download|postOriginal|upload)\s*(?:<[^>]*>)?\s*\(\s*\{([\s\S]*?)\}\s*\)/g;
  let m;
  while ((m = regex.exec(content)) !== null) {
    const method = m[1].toLowerCase() === 'postoriginal' ? 'POST' : m[1].toUpperCase();
    const body = m[2];
    const urlMatch = body.match(/url:\s*([`'"][^`'"]+[`'"])/);
    if (urlMatch) {
      let rawUrl = urlMatch[1].slice(1, -1);
      // strip template literal interpolations for path key
      const cleanUrl = rawUrl.split('?')[0].replace(/\$\{[^}]+\}/g, '{param}');
      const key = `${method} ${cleanUrl}`;
      if (!adminEndpoints.has(key)) {
        adminEndpoints.set(key, {
          app: 'admin',
          prefix: '/admin-api',
          method,
          url: cleanUrl,
          rawUrl,
          files: new Set()
        });
      }
      adminEndpoints.get(key).files.add(file.replace(/^apps\//, ''));
    }
  }
}

// Parse miniapp API files and components
for (const file of miniappFiles) {
  const content = fs.readFileSync(file, 'utf8');
  // Find request({ url: ..., method: ... })
  const regex = /request\s*\(\s*\{([\s\S]*?)\}\s*\)/g;
  let m;
  while ((m = regex.exec(content)) !== null) {
    const body = m[1];
    const urlMatch = body.match(/url:\s*([`'"][^`'"]+[`'"])/);
    if (urlMatch) {
      const methodMatch = body.match(/method:\s*([`'"][^`'"]+[`'"])/);
      const method = methodMatch ? methodMatch[1].slice(1, -1).toUpperCase() : 'GET';
      let rawUrl = urlMatch[1].slice(1, -1);
      const cleanUrl = rawUrl.split('?')[0].replace(/\$\{[^}]+\}/g, '{param}');
      const key = `${method} ${cleanUrl}`;
      if (!miniappEndpoints.has(key)) {
        miniappEndpoints.set(key, {
          app: 'miniapp',
          prefix: '/app-api',
          method,
          url: cleanUrl,
          rawUrl,
          files: new Set()
        });
      }
      miniappEndpoints.get(key).files.add(file.replace(/^apps\//, ''));
    }
  }
}

console.log(`Admin unique endpoints: ${adminEndpoints.size}`);
console.log(`Miniapp unique endpoints: ${miniappEndpoints.size}`);

const allAdmin = Array.from(adminEndpoints.values()).map(x => ({ ...x, files: Array.from(x.files) }));
const allMiniapp = Array.from(miniappEndpoints.values()).map(x => ({ ...x, files: Array.from(x.files) }));

fs.writeFileSync('all_endpoints.json', JSON.stringify({ admin: allAdmin, miniapp: allMiniapp }, null, 2));
