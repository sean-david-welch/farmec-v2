import { FastifyReply, FastifyRequest } from 'fastify';
import SupplierService from 'services/suppliers';

class SuppliersController {
  private supplierService: SupplierService;

  constructor(supplierService: SupplierService) {
    this.supplierService = supplierService;
  }

  async GetSuppliers(request: FastifyRequest, reply: FastifyReply) {
    try {
      const suppliers = await this.supplierService.GetSuppliers();
      return reply.code(200).send(suppliers);
    } catch (error) {
      console.error(error);
      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }
}

export default SuppliersController;
