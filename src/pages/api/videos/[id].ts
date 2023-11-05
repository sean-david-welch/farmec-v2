import type { APIRoute } from 'astro';

import { pool } from '../../../database/connection';
import { youtube as YouTube } from '@googleapis/youtube';

import type { YoutubeApiResponse } from '../../../types/video';

export const GET: APIRoute = async ({ params }) => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const query = await client.query(
      'SELECT * FROM "Video" WHERE "supplierId" = $1;',
      [id]
    );
    return new Response(JSON.stringify(query.rows), {
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    console.error(error);
    return new Response(JSON.stringify({ error: 'Database query failed' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  } finally {
    client.release();
  }
};

export const PUT: APIRoute = async ({ request, params }): Promise<Response> => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const data = await request.json();

    const { web_url } = data;

    const videoId = web_url.split('v=')[1].split('&')[0];

    const youtube = YouTube({
      version: 'v3',
      auth: process.env.YOUTUBE_API_KEY,
    });

    const videoResponse = (await youtube.videos.list({
      part: ['id', 'snippet'],
      id: videoId,
      maxResults: 1,
    })) as YoutubeApiResponse;

    if (!videoResponse.data.items || videoResponse.data.items.length === 0) {
      throw new Error('Video not found on YouTube');
    }

    const { title, thumbnails, description } =
      videoResponse.data.items[0].snippet;

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
      data.supplierId,
      web_url,
      title,
      description,
      videoResponse.data.items[0].id,
      thumbnails.medium.url,
    ];

    const result = await client.query(sql, values);

    return new Response(JSON.stringify(result.rows[0]), {
      headers: { 'Content-Type': 'application/json' },
    });
  } catch (error) {
    console.error(error);
    return new Response(JSON.stringify({ error: 'Database query failed' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  } finally {
    client.release();
  }
};

export const DELETE: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const query = await client.query(
      'DELETE FROM Video WHERE id = $1 RETURNING *;',
      [id]
    );

    if (query.rowCount === 0) {
      return new Response(
        JSON.stringify({ message: 'No Video found with that ID.' }),
        {
          status: 404,
          headers: { 'Content-Type': 'application/json' },
        }
      );
    }

    return new Response(
      JSON.stringify({ message: 'Video deleted successfully.' }),
      {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }
    );
  } catch (error) {
    console.error(error);
    return new Response(
      JSON.stringify({ error: 'Error performing the delete operation' }),
      {
        status: 500,
        headers: { 'Content-Type': 'application/json' },
      }
    );
  } finally {
    client.release();
  }
};
