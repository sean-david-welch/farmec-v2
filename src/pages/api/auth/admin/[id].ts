import type { APIRoute } from 'astro';

import { pool } from '../../../../database/connection';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const query = await client.query('SELECT * FROM "users" WHERE "uid" = $1;', [id]);
    const user = query.rows[0];

    if (!user) {
      return new Response(JSON.stringify({ error: 'User not found' }), {
        status: 404,
        headers: { 'Content-Type': 'application/json' },
      });
    }

    const response = {
      isAdmin: user.role === 'admin',
      userDetails: user.role === 'admin' ? user : { uid: user.uid, role: user.role },
    };

    return new Response(JSON.stringify(response), {
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
