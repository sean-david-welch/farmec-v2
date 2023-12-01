import type { APIRoute } from 'astro';

import { getAuth } from 'firebase-admin/auth';
import { pool } from '../../../database/connection';

export const POST: APIRoute = async ({ request }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const data = await request.json();

    const { email, password, role } = data;

    if (!email || !password) {
      return new Response('Missing form data', { status: 400 });
    }

    const userRecord = await getAuth().createUser({ email, password });

    const sql = `INSERT INTO users(uid, email, role) VALUES($1, $2, $3) RETURNING *`;

    const values = [userRecord.uid, email, role];

    const result = await client.query(sql, values);

    return new Response(JSON.stringify(result.rows[0]), {
      status: 200,
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
