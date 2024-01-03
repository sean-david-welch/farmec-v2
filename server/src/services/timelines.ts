import { FastifyInstance } from 'fastify';
import { Timeline } from 'types/about';

import { v4 as uuidv4 } from 'uuid';

class TimelineService {
  fastify: FastifyInstance;

  constructor(fastify: FastifyInstance) {
    this.fastify = fastify;
  }

  async getTimelines() {
    const client = await this.fastify.pg.connect();

    try {
      const query = await client.query('select * from "Timeline"');

      return query.rows;
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }

  async createTimeline(data: Timeline) {
    const client = await this.fastify.pg.connect();

    const uuid = uuidv4();
    const currentDateTime = new Date().toISOString();

    try {
      const sql = `INSERT INTO Timeline(id, title, date, body, created)
      VALUES($1, $2, $3, $4, $5)
      RETURNING *;
      `;

      const values = [uuid, data.title, data.date, data.body, currentDateTime];

      const result = await client.query(sql, values);

      return result.rows[0];
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }

  async updateTimeline(id: string, data: Timeline) {}

  async deleteTimeline(id: string) {}
}

export default TimelineService;
