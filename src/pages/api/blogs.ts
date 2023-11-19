import { v4 as uuidv4 } from 'uuid';
import type { APIRoute } from 'astro';
import { verifyToken } from '../../utils/admin';

import { pool } from '../../database/connection';
import { generatePresignedUrl } from '../../utils/aws';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query('select * from "Blog"');
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

export const POST: APIRoute = async ({ request }): Promise<Response> => {
  const token = request.headers.get('Authorization')?.split('Bearer ')[1];
  if (!token) {
    return new Response(JSON.stringify({ error: 'No token provided' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' },
    });
  }

  await verifyToken(token);
  const client = await pool.connect();

  try {
    const uuid = uuidv4();
    const data = await request.json();
    const currentDateTime = new Date().toISOString();

    const sql = `INSERT INTO "Blog"(id, title, date, main_image, subheading, body, created)
    VALUES($1, $2, $3, $4, $5, $6, $7)
    RETURNING *;
    `;

    const values = [
      uuid,
      data.title,
      data.date,
      data.main_image as string,
      data.subheading,
      data.body,
      currentDateTime,
    ];

    const result = await client.query(sql, values);

    const presignedUrl = await generatePresignedUrl('farmec-bucket', data.main_image as string);

    console.log(presignedUrl);

    return new Response(JSON.stringify({ data: result.rows[0], uploadUrl: presignedUrl }), {
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
