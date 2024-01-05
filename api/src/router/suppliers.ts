import { FastifyInstance } from 'fastify';
import SupplierService from '../services/suppliers';
import SuppliersController from '../controllers/suppliers';

const suppliers = async (fastify: FastifyInstance) => {
  const supplierService = new SupplierService(fastify);
  const suppliersController = new SuppliersController(supplierService);

  fastify.get('/suppliers', suppliersController.getSuppliers.bind(suppliersController));
  fastify.post('/suppliers', suppliersController.createSupplier.bind(suppliersController));

  fastify.get('/suppliers/:id', suppliersController.getSuppliers.bind(suppliersController));
  fastify.put('/suppliers/:id', suppliersController.updateSupplier.bind(suppliersController));
  fastify.delete('/suppliers/:id', suppliersController.deleteSupplier.bind(suppliersController));
};

export default suppliers;
