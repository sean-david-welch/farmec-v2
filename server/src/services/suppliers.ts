import { FastifyInstance } from 'fastify';

class SupplierService {
  fastify: FastifyInstance;

  constructor(fastify: FastifyInstance) {
    this.fastify = fastify;
  }

  async getSuppliers() {
    const client = await this.fastify.pg.connect();
    try {
      const query = await client.query('select * from "Supplier"');
      console.log('Service: Query executed', query.rows);

      return query.rows;
    } catch (error) {
      console.error('Service Error:', error);
      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }
}

export default SupplierService;
