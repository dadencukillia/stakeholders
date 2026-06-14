// @ts-check

import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "astro/config";
import react from "@astrojs/react";
import path from 'node:path';

// https://astro.build/config
export default defineConfig({
	vite: {
		plugins: [tailwindcss()],
    resolve: {
      alias: {
        '@': path.resolve(import.meta.dirname, './src')
      }
    }
	},
	integrations: [react()],
	output: "server",
});
