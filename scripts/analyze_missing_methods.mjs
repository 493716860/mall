import fs from 'node:fs';

const missingData = JSON.parse(fs.readFileSync('/tmp/exact_missing.json', 'utf8'));
const allMissing = [...missingData.adminMissing, ...missingData.miniappMissing];

console.log('Total missing items to resolve:', allMissing.length);

const routeDetails = [];
for (const item of allMissing) {
  const content = fs.readFileSync(item.file, 'utf8');
  let cleanPath = item.api.replace('/admin-api', '').replace('/app-api', '');
  
  let method = 'GET';
  // Common keywords in ruoyi-vue-pro
  if (cleanPath.includes('/create') || cleanPath.includes('/add') || cleanPath.includes('/submit') || cleanPath.includes('/take') || cleanPath.includes('/login') || cleanPath.includes('/send') || cleanPath.includes('/bind') || cleanPath.includes('/upload')) {
    method = 'POST';
  } else if (cleanPath.includes('/update') || cleanPath.includes('/receive') || cleanPath.includes('/delivery') || cleanPath.includes('/agree') || cleanPath.includes('/disagree') || cleanPath.includes('/refuse') || cleanPath.includes('/approve') || cleanPath.includes('/reject') || cleanPath.includes('/close') || cleanPath.includes('/use')) {
    method = 'PUT';
  } else if (cleanPath.includes('/delete') || cleanPath.includes('/clean') || cleanPath.includes('/cancel') || cleanPath.includes('/unbind')) {
    method = 'DELETE';
  }

  routeDetails.push({ api: item.api, method, file: item.file });
}

fs.writeFileSync('/tmp/missing_with_methods.json', JSON.stringify(routeDetails, null, 2));
console.log('Saved detailed method information to /tmp/missing_with_methods.json');
