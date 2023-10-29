import type { APIRoute } from 'astro';
import { pool } from '../../../database/connection';

export const GET: APIRoute = async ({ params }) => {
  const client = await pool.connect();
  const supplierId = params.supplierId;

  try {
    const query = await client.query(
      'SELECT * FROM "Machine" WHERE "supplierId" = $1;',
      [supplierId]
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
