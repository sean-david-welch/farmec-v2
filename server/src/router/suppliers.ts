const routes = async (fastify, options) => {
  fastify.get('/suppliers', getSuppliersController);
};
