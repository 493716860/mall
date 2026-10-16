import { defineConfig, loadEnv, normalizePath } from 'vite'
import { resolve, relative, dirname } from 'node:path'
import Vue from '@vitejs/plugin-vue'
import VueJsx from '@vitejs/plugin-vue-jsx'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import UnoCSS from 'unocss/vite'
import { fileURLToPath } from 'node:url'
const adminDir = dirname(fileURLToPath(import.meta.url))

export default defineConfig(({ mode }) => {
  const env = {
    ...loadEnv(mode, resolve(adminDir, '../..')),
    ...loadEnv(mode, adminDir)
  }
  const isDemoMode = env.VITE_DEMO_MODE !== 'false'
  return {
    envDir: adminDir,
    base: env.VITE_BASE_PATH || '/',
    define: {
      'import.meta.env.VITE_DEMO_MODE': JSON.stringify(isDemoMode ? 'true' : 'false')
    },
    plugins: [
      Vue(),
      VueJsx(),
      UnoCSS(),
      AutoImport({
        imports: [
          'vue',
          'vue-router',
          {
            '@/hooks/web/useI18n': ['useI18n'],
            '@/hooks/web/useMessage': ['useMessage'],
            '@/hooks/web/useTable': ['useTable'],
            '@/hooks/web/useCrudSchemas': ['useCrudSchemas'],
            '@/utils/formRules': ['required'],
            '@/utils/dict': ['DICT_TYPE']
          }
        ],
        dts: 'src/types/auto-imports.d.ts',
        resolvers: [ElementPlusResolver()]
      }),
      Components({
        dts: 'src/types/auto-components.d.ts',
        resolvers: [ElementPlusResolver()],
        globs: ['src/components/**/*.vue', '!src/components/DiyEditor/components/mobile/**']
      })
    ],
    resolve: { alias: { '@': resolve(process.cwd(), 'src') } },
    server: {
      host: '0.0.0.0',
      port: 3000,
      proxy: {
        '/admin-api': {
          target: 'http://127.0.0.1:48080',
          changeOrigin: true
        },
        '/app-api': {
          target: 'http://127.0.0.1:48080',
          changeOrigin: true
        }
      }
    },
    preview: { host: '0.0.0.0', port: 3000 },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: (source, filename) => {
            if (filename.endsWith('variables.scss') || filename.endsWith('global.module.scss'))
              return source
            const path = normalizePath(
              relative(
                dirname(filename.split('?')[0]),
                resolve(process.cwd(), 'src/styles/variables.scss')
              )
            )
            return `@use "${path.startsWith('.') ? path : './' + path}" as *;\n${source}`
          }
        }
      }
    },
    build: { outDir: 'dist', chunkSizeWarningLimit: 2000, sourcemap: false }
  }
})
