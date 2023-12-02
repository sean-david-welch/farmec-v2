import type { APIRoute } from 'astro';

import { getAuth } from 'firebase-admin/auth';
import { pool } from '../../../database/connection';

export const PUT: APIRoute = async ({ request, params }): Promise<Response> => {
  const client = await pool.connect();
  const id = params.id as string;

  try {
    const data = await request.json();
    const { email, password, role } = data;

    const userRecord = await getAuth().updateUser(id, { email: email, password: password });

    const sql = `INSERT INTO users(uid, email, role) VALUES($1, $2, $3) RETURNING *`;

    const values = [userRecord.uid, email, role];

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
  const id = params.id as string;

  try {
    await getAuth().deleteUser(id);
    const query = await client.query('DELETE FROM users WHERE id = $1 RETURNING *;', [id]);

    if (query.rowCount === 0) {
      return new Response(JSON.stringify({ message: 'No User found with that ID.' }), {
        status: 404,
        headers: { 'Content-Type': 'application/json' },
      });
    }

    return new Response(JSON.stringify({ message: 'User deleted successfully.' }), {
      status: 200,
      headers: { 'Content-Type': 'application/json' },
    });
  } catch (error) {
    console.error(error);
    return new Response(JSON.stringify({ error: 'Error performing the delete operation' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  } finally {
    client.release();
  }
};
