import { FastifyReply, FastifyRequest } from 'fastify';
import SupplierService from 'services/suppliers';

class SuppliersController {
  private supplierService: SupplierService;

  constructor(supplierService: SupplierService) {
    this.supplierService = supplierService;
  }

  async getSuppliers(request: FastifyRequest, reply: FastifyReply) {
    try {
      const suppliers = await this.supplierService.getSuppliers();
      return reply.code(200).send(suppliers);
    } catch (error) {
      console.error('Controller Error:', error);
      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }
}

export default SuppliersController;
