import { defineConfig } from "astro/config";
import tailwind from "@astrojs/tailwind";

export default defineConfig({
  // Wails embeds the built dist/ folder – output must be static
  output: "static",

  // During `wails dev`, Wails proxies to the Astro dev server on port 4321
  server: {
    port: 4321,
    host: true,
  },

  integrations: [tailwind()],
});