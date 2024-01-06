import { Video, VideoDTO } from 'types/video';
import { v4 as uuidv4 } from 'uuid';
import { FastifyInstance } from 'fastify';

class VideoService {
  fastify: FastifyInstance;

  constructor(fastify: FastifyInstance) {
    this.fastify = fastify;
  }

  async getVideos(id: string) {
    const client = await this.fastify.pg.connect();
    try {
      const query = await client.query('SELECT * FROM "Video" WHERE "supplierId" = $1;', [id]);

      return query.rows;
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }

  async createVideo(data: VideoDTO) {
    const client = await this.fastify.pg.connect();

    const uuid = uuidv4();
    const currentDateTime = new Date().toISOString();

    try {
      const sql = `INSERT INTO Video(id, supplierId, web_url, title, description, video_id, thumbnail_url, created)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING *;
        `;

      const values = [
        uuid,
        data.video.supplierId,
        data.video.web_url,
        data.youtube.title,
        data.youtube.description,
        data.youtube.id,
        data.youtube.thumbnail,
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

  async updateVideo(id: string, data: VideoDTO) {
    const client = await this.fastify.pg.connect();

    try {
      const sql = `UPDATE Machine SET
        supplierId = $2,
        web_url = $3,
        title = $4,
        description = $5,
        video_id = $6,
        thumbnail_url = $7,
        WHERE id = $1
        RETURNING *;
        `;

      const values = [
        id,
        data.video.supplierId,
        data.video.web_url,
        data.youtube.title,
        data.youtube.description,
        data.youtube.id,
        data.youtube.thumbnail,
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

  async deleteVideo(id: string) {
    const client = await this.fastify.pg.connect();

    try {
      const query = await client.query('DELETE FROM Video WHERE id = $1 RETURNING *;', [id]);

      return query.rows[0];
    } catch (error) {
      console.error('Service Error:', error);

      throw new Error('Database query failed');
    } finally {
      client.release();
    }
  }
}

export default VideoService;
