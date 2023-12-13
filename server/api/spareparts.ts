import type { APIRoute } from 'astro';
import { pool } from '../../database/connection';

import { v4 as uuidv4 } from 'uuid';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query('select * from "SpareParts"');
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
  const client = await pool.connect();

  const uuid = uuidv4();

  try {
    const data = await request.json();

    const sql = `INSERT INTO SpareParts(id, supplierId, name, parts_image, spare_parts_link, pdf_link)
    VALUES($1, $2, $3, $4, $5, $6)
    RETURNING *;
    `;

    const values = [
      uuid,
      data.supplierId,
      data.name,
      data.parts_image,
      data.spare_parts_link,
      data.pdf_link,
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
