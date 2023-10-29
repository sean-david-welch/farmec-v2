import { pool } from '../../database/connection';
import type { APIRoute } from 'astro';

import { v4 as uuidv4 } from 'uuid';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query('select * from "Product"');
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

  try {
    const data = await request.json();
    const uuid = uuidv4();

    const sql = `INSERT INTO Product(id, machineId,  name, product_image, description, product_link)
    VALUES($1, $2, $3, $4, $5, $6)
    RETURNING *;
    `;

    const values = [
      uuid,
      data.machineId,
      data.name,
      data.machine_image,
      data.description,
      data.machine_link,
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
