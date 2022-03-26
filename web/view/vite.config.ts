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
  }
})
