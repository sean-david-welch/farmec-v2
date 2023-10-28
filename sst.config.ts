import type { SSTConfig } from 'sst';
import { AstroSite } from 'sst/constructs';

export default {
  config(_input: any) {
    return {
      name: 'farmec-astro',
      region: 'eu-west-1',
    };
  },
  stacks(app: any) {
    app.stack(function Site({ stack }: any) {
      const site = new AstroSite(stack, 'site');
      stack.addOutputs({
        url: site.url,
      });
    });
  },
} satisfies SSTConfig;
