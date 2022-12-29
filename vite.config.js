import { fileURLToPath, URL } from "node:url";
import * as dotenv from "dotenv";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

dotenv.config();

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    port: process.env.CLIENT_PORT,
  }
});
