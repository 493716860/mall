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

const javaFiles = walk('apps/server-java', ['.java']);
let javaApiCount = 0;
const modulesCount = {};

for (const file of javaFiles) {
  const content = fs.readFileSync(file, 'utf8');
  const matches = content.matchAll(/@(Get|Post|Put|Delete|Request)Mapping/g);
  for (const m of matches) {
    javaApiCount++;
    const parts = file.split(path.sep);
    const mod = parts[2] || 'other';
    modulesCount[mod] = (modulesCount[mod] || 0) + 1;
  }
}

console.log('Total Java Controller Mapping annotations:', javaApiCount);
console.log('Java Module Breakdown:', modulesCount);
