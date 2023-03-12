import { fileURLToPath, URL } from "node:url";
import * as dotenv from "dotenv";
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

dotenv.config();

// https://vitejs.dev/config/
export default defineConfig({
    define: {
        BUILD_API_URL: JSON.stringify(process.env.SERVER_API_URL)
    },
    plugins: [
        vue(),
    ],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./client", import.meta.url)),
        },
    },
    server: {
        port: process.env.CLIENT_PORT,
    }
});
