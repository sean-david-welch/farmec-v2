import { FastifyInstance } from 'fastify';

class SupplierService {
  fastify: FastifyInstance;

  constructor(fastify: FastifyInstance) {
    this.fastify = fastify;
  }

  async GetSuppliers() {
    const client = await this.fastify.pg.connect();
    try {
      const query = await client.query('select * from "Supplier"');
      return query.rows;
    } catch (error) {
      console.error(error);
      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }
}

export default SupplierService;
