import { pool } from '../../database/connection';

export const GET = async () => {
  const client = await pool.connect();
  try {
    const query = await client.query('select * from suppliers');
    return new Response(JSON.stringify(query.rows), {
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
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

export const POST = async () => {};
