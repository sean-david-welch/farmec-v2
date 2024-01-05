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

  async getSupplierId(id: string) {
    const client = await this.fastify.pg.connect();

    try {
      const query = await client.query('select * from "Supplier" where id = $1', [id]);

      return query.rows[0];
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

  async updateSupplier(id: string, data: Supplier) {
    const client = await this.fastify.pg.connect();

    try {
      const sql = `UPDATE Supplier SET
      name = $2,
      logo_image = $3,
      marketing_image = $4,
      description = $5,
      social_facebook = $6,
      social_twitter = $7,
      social_instagram = $8,
      social_youtube = $9,
      social_linkedin = $10,
      social_website = $11,
      created = $12
      WHERE id = $1
      RETURNING *;`;

      const values = [
        id,
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

  async deleteSupplier(id: string) {
    const client = await this.fastify.pg.connect();

    try {
      const query = await client.query('DELETE FROM Supplier WHERE id = $1 RETURNING *;', [id]);

      return query.rows[0];
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }
}

export default SupplierService;
