import { fileURLToPath, URL } from "node:url";
import * as dotenv from "dotenv";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

dotenv.config();

// https://github.com/vuetifyjs/vuetify-loader/tree/next/packages/vite-plugin
import vuetify from 'vite-plugin-vuetify'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
		vue(),
		vuetify({ autoImport: true }),
	],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  server: {
    port: process.env.CLIENT_PORT,
  }
});
