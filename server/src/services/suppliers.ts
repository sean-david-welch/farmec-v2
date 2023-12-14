import { FastifyInstance } from 'fastify';

import { v4 as uuidv4 } from 'uuid';

import Supplier from 'types/supplier';

class SupplierService {
  fastify: FastifyInstance;

  constructor(fastify: FastifyInstance) {
    this.fastify = fastify;
  }

  async getSuppliers() {
    const client = await this.fastify.pg.connect();

    try {
      const query = await client.query('select * from "Supplier"');

      return query.rows;
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }

  async createSupplier(data: Supplier) {
    const client = await this.fastify.pg.connect();

    const uuid = uuidv4();
    const currentDateTime = new Date().toISOString();

    try {
      const sql = `INSERT INTO Supplier(...) VALUES($1, $2, ..., $12) RETURNING *;`;

      const values = [
        uuid,
        data.name,
        data.logo_image,
        data.marketing_image,
        data.description,
        data.social_facebook,
        data.social_twitter,
        data.social_instagram,
        data.social_youtube,
        data.social_linkedin,
        data.social_website,
        currentDateTime,
      ];

      const result = await client.query(sql, values);
      return result.rows[0];
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }
}

export default SupplierService;
