import { pool } from '../../database/connection';
import { v4 as uuidv4 } from 'uuid';

import type { APIRoute } from 'astro';
import type { YoutubeApiResponse } from '../../types/video';

import { youtube as YouTube } from '@googleapis/youtube';

export const POST: APIRoute = async ({ request }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const data = await request.json();
    const uuid = uuidv4();

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

    const sql = `INSERT INTO Video(id, supplierId, web_url, title, description, video_id, thumbnail_url)
    VALUES($1, $2, $3, $4, $5, $6, $7)
    RETURNING *;
    `;

    const values = [
      uuid,
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
