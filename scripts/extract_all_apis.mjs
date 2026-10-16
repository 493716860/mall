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
    } else if (file.endsWith('.ts') || file.endsWith('.js')) {
      results.push(full);
    }
  }
  return results;
}

function extractAdminApis() {
  const files = walk('apps/admin/src/api');
  const apis = [];
  for (const file of files) {
    const content = fs.readFileSync(file, 'utf8');
    const lines = content.split('\n');
    let currentExport = '';
    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      const fnMatch = line.match(/export\s+(?:const|function)\s+([a-zA-Z0-9_$]+)/);
      if (fnMatch) currentExport = fnMatch[1];
      
      const reqMatch = line.match(/request\.(get|post|put|delete|download|postOriginal|upload)\s*\(\s*\{/);
      if (reqMatch) {
        const method = reqMatch[1].toUpperCase();
        // find url in following lines
        let url = '';
        for (let j = i; j < Math.min(i + 15, lines.length); j++) {
          const urlMatch = lines[j].match(/url:\s*([`'"][^`'"]+[`'"])/);
          if (urlMatch) {
            url = urlMatch[1].slice(1, -1);
            break;
          }
        }
        if (url) {
          apis.push({
            app: 'admin',
            method: method === 'POSTORIGINAL' ? 'POST' : method,
            functionName: currentExport,
            url: url.startsWith('/') ? url : '/' + url,
            file: file.replace(/^apps\//, '')
          });
        }
      }
    }
  }
  return apis;
}

function extractMiniappApis() {
  const files = walk('apps/miniapp/sheep/api');
  const apis = [];
  for (const file of files) {
    const content = fs.readFileSync(file, 'utf8');
    const lines = content.split('\n');
    let currentMethodName = '';
    let currentComment = '';
    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      const commentMatch = line.match(/\/\/\s*(.*)/);
      if (commentMatch) {
        currentComment = commentMatch[1].trim();
      }
      const fnMatch = line.match(/^\s*([a-zA-Z0-9_$]+)\s*:\s*(?:\([^)]*\)|[a-zA-Z0-9_$]+)\s*=>/);
      if (fnMatch) {
        currentMethodName = fnMatch[1];
      }
      if (line.includes('return request({') || line.includes('request({')) {
        let url = '';
        let method = 'GET';
        for (let j = i; j < Math.min(i + 15, lines.length); j++) {
          const urlMatch = lines[j].match(/url:\s*([`'"][^`'"]+[`'"])/);
          if (urlMatch) {
            url = urlMatch[1].slice(1, -1);
          }
          const methodMatch = lines[j].match(/method:\s*([`'"][^`'"]+[`'"])/);
          if (methodMatch) {
            method = methodMatch[1].slice(1, -1).toUpperCase();
          }
          if (lines[j].includes('})')) break;
        }
        if (url) {
          apis.push({
            app: 'miniapp',
            method,
            functionName: currentMethodName,
            comment: currentComment,
            url: url.startsWith('/') ? url : '/' + url,
            file: file.replace(/^apps\//, '')
          });
          currentComment = '';
        }
      }
    }
  }
  return apis;
}

const adminApis = extractAdminApis();
const miniappApis = extractMiniappApis();

console.log(`Admin APIs: ${adminApis.length}`);
console.log(`Miniapp APIs: ${miniappApis.length}`);

fs.writeFileSync('extracted_apis.json', JSON.stringify({ adminApis, miniappApis }, null, 2));
