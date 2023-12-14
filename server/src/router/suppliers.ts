import { FastifyInstance } from 'fastify';
import SupplierService from '../services/suppliers';
import SuppliersController from '../controllers/suppliers';

const suppliers = async (fastify: FastifyInstance) => {
  const supplierService = new SupplierService(fastify);
  const suppliersController = new SuppliersController(supplierService);

  fastify.get('/suppliers', suppliersController.getSuppliers.bind(suppliersController));
  fastify.post('/suppliers', suppliersController.createSupplier.bind(suppliersController));

  fastify.put('/suppliers', suppliersController.updateSupplier.bind(suppliersController));
};

export default suppliers;
