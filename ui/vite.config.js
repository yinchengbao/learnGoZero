import { fileURLToPath, URL } from 'node:url'

import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import importPlugin from '@opentiny/vue-vite-import'


export default defineConfig({
  plugins: [
    vue(),
    importPlugin(
        [
          {
            libraryName: '@opentiny/vue'
          },
          {
            libraryName: `@opentiny/vue-icon`,
            customName: (name) => {
              return `@opentiny/vue-icon/lib/${name.replace(/^icon-/, '')}.js`
            }
          }
        ], 'pc'
    )
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: ['.js', '.jsx', '.vue']
  },
  define: {
      'process.env': Object.assign({}, process.env)
  },
   server:{
        proxy: {// 跨域代理
            '/apis': {
                target: 'http://localhost:1012', //
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/apis/, '')
            },
        }
    },
})
