import { FastifyInstance, FastifyPluginOptions } from 'fastify';

import supplierRoutes from './suppliers';

const mainRouter = async (fastify: FastifyInstance, options: FastifyPluginOptions) => {
  fastify.register(supplierRoutes);
};

export default mainRouter;
