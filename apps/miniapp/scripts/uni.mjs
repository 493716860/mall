import { spawn } from 'node:child_process';
import { fileURLToPath } from 'node:url';
import { dirname, resolve } from 'node:path';
const root = resolve(dirname(fileURLToPath(import.meta.url)), '..');
const child = spawn(
  process.execPath,
  [resolve(root, 'node_modules/@dcloudio/vite-plugin-uni/bin/uni.js'), ...process.argv.slice(2)],
  { cwd: root, env: { ...process.env, UNI_INPUT_DIR: root, CI: '1' }, stdio: 'inherit' },
);
child.on('error', (error) => {
  console.error(error.message);
  process.exitCode = 1;
});
child.on('exit', (code) => {
  process.exitCode = code ?? 1;
});
