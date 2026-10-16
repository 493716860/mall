import { spawn } from 'node:child_process'
import { fileURLToPath } from 'node:url'
import { dirname, resolve } from 'node:path'

const rootDir = resolve(dirname(fileURLToPath(import.meta.url)), '..')

console.log('[DevRunner] Starting Admin Web server on port 3000...')
const admin = spawn('npx', ['vite', '--mode', 'demo', '--host', '0.0.0.0', '--port', '3000'], {
  cwd: resolve(rootDir, 'apps/admin'),
  stdio: 'inherit',
  env: { ...process.env }
})

const handleSignal = () => {
  try {
    admin.kill()
  } catch (e) {}
  process.exit(0)
}

process.on('SIGINT', handleSignal)
process.on('SIGTERM', handleSignal)
