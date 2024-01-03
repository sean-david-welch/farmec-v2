import { FastifyInstance, FastifyPluginOptions } from 'fastify';

import suppliers from './suppliers';
import videos from './videos';

const mainRouter = async (fastify: FastifyInstance, options: FastifyPluginOptions) => {
  fastify.register(suppliers);
  fastify.register(videos);
};

export default mainRouter;
