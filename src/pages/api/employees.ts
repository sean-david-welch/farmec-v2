import { v4 as uuidv4 } from 'uuid';
import type { APIRoute } from 'astro';
import { pool } from '../../database/connection';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query('select * from "Employee"');
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

    const sql = `INSERT INTO Employee(id, name, email, role, bio, profile_image, created, phone)
      VALUES($1, $2, $3, $4, $5, $6, $7, $8)
      RETURNING *;
      `;

    const uuid = uuidv4();
    const currentDateTime = new Date().toISOString();

    const values = [
      uuid,
      data.name,
      data.email,
      data.role,
      data.bio,
      data.profile_image,
      currentDateTime,
      data.phone,
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
