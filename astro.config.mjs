import { defineConfig } from 'astro/config';
import solidJs from '@astrojs/solid-js';
import tailwind from '@astrojs/tailwind';

import deno from '@astrojs/deno';

export default defineConfig({
  output: 'server',
  adapter: deno(),
  integrations: [solidJs(), tailwind()],
});
