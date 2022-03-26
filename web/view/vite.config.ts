import path from 'path'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  // https://vitejs.dev/guide/backend-integration.html
  build: {
    manifest: true,
    rollupOptions: {
      input: './src/main.ts'
    }
  },
  // https://www.scien.cx/2021/11/04/how-to-configure-import-aliases-in-vite-typescript-and-jest/
  resolve: {
    alias: {
      '@root': path.resolve(__dirname, '../..')
    }
  }
})
