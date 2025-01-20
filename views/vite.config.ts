import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  appType: 'spa',
  css: {
    transformer: 'postcss',
  },
  optimizeDeps: {
    esbuildOptions: {
      minify: true,
      minifyWhitespace: true,
      minifyIdentifiers: true,
      minifySyntax: true,
      keepNames: false,
    },
  }
})
