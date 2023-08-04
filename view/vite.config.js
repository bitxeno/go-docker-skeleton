import { defineConfig } from "vite";
import { resolve } from "path";
import vue from "@vitejs/plugin-vue";
import Icons from 'unplugin-icons/vite'
import svgLoader from "vite-svg-loader";

export default defineConfig({
  plugins: [vue(), Icons(), svgLoader()],
  resolve: {
    alias: [
      {
        find: "@",
        replacement: resolve(__dirname, "src"),
      },
    ],
  },
});
