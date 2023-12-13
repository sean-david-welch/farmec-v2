import { FastifyInstance, FastifyPluginOptions } from 'fastify';

import suppliers from './suppliers';

const mainRouter = async (fastify: FastifyInstance, options: FastifyPluginOptions) => {
  fastify.register(suppliers);
};

export default mainRouter;
