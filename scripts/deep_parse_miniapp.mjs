import fs from 'node:fs';
import path from 'node:path';

function walk(dir) {
  let results = [];
  if (!fs.existsSync(dir)) return results;
  const list = fs.readdirSync(dir);
  for (const f of list) {
    const full = path.join(dir, f);
    const stat = fs.statSync(full);
    if (stat.isDirectory()) {
      results = results.concat(walk(full));
    } else if (f.endsWith('.js') && !f.endsWith('index.js')) {
      results.push(full);
    }
  }
  return results;
}

const apiFiles = walk('apps/miniapp/sheep/api');
const allMethods = [];

for (const file of apiFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const relPath = file.replace('apps/miniapp/', '');
  
  const lines = content.split('\n');
  let currentComment = '';
  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    const cMatch = line.match(/\/\/\s*(.+)/);
    if (cMatch) {
      currentComment = cMatch[1].trim();
    }
    
    const mMatch = line.match(/^\s*([a-zA-Z0-9_$]+)\s*[:(]/);
    if (mMatch) {
      const name = mMatch[1];
      if (['return', 'import', 'export', 'const', 'let', 'var', 'custom', 'params', 'data', 'url', 'method', 'header'].includes(name)) continue;
      
      let url = '';
      let httpMethod = 'GET';
      for (let j = i; j < Math.min(i + 25, lines.length); j++) {
        const uMatch = lines[j].match(/url:\s*([`'"][^`'"]+[`'"])/);
        if (uMatch) {
          url = uMatch[1].slice(1, -1);
        }
        const methMatch = lines[j].match(/method:\s*([`'"][^`'"]+[`'"])/);
        if (methMatch) {
          httpMethod = methMatch[1].slice(1, -1).toUpperCase();
        }
        if (lines[j].includes('},') || lines[j].trim() === '};' || (lines[j].trim() === '}' && lines[j+1]?.includes(':'))) {
          if (url) break;
        }
      }
      
      if (url) {
        allMethods.push({
          file: relPath,
          methodName: name,
          comment: currentComment || name,
          url: url,
          httpMethod: httpMethod
        });
      }
      currentComment = '';
    }
  }
}

console.log('Total API methods with URL in sheep/api:', allMethods.length);
const uniqueUrls = new Set(allMethods.map(m => `${m.httpMethod} ${m.url.split('?')[0]}`));
console.log('Unique endpoints in miniapp:', uniqueUrls.size);

// Group by category
const byModule = {};
for (const m of allMethods) {
  const mod = m.file.split('/')[2];
  byModule[mod] = (byModule[mod] || 0) + 1;
}
console.log('API methods by module:', byModule);

fs.writeFileSync('miniapp_all_api_methods.json', JSON.stringify(allMethods, null, 2));
