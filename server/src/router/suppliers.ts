import { FastifyInstance } from 'fastify';
import SupplierService from '../services/suppliers';
import SuppliersController from '../controllers/suppliers';

const routes = async (fastify: FastifyInstance) => {
  const supplierService = new SupplierService(fastify);
  const suppliersController = new SuppliersController(supplierService);

  fastify.get('/suppliers', suppliersController.GetSuppliers.bind(suppliersController));
};

export default routes;
