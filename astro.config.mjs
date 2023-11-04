import { defineConfig } from 'astro/config';

import aws from 'astro-sst/lambda';
import solidJs from '@astrojs/solid-js';
import tailwind from '@astrojs/tailwind';

export default defineConfig({
  output: 'server',
  adapter: aws(),
  integrations: [solidJs(), tailwind()],
});
