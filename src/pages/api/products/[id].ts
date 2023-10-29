import type { APIRoute } from 'astro';
import { pool } from '../../../database/connection';

export const GET: APIRoute = async ({ params }) => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const query = await client.query(
      'SELECT * FROM "Machine" WHERE "supplierId" = $1;',
      [id]
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

export const PUT: APIRoute = async ({ request, params }): Promise<Response> => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const data = await request.json();

    const sql = `UPDATE Product SET
    machineId = $2,
    name = $3,
    product_image = $4,
    description = $5,
    product_link = $6,
    WHERE id = $1
    RETURNING *;`;

    const values = [
      id,
      data.supplierId,
      data.name,
      data.product_image,
      data.description,
      data.product_link,
    ];

    const result = await client.query(sql, values);

    if (result.rows.length > 0) {
      return new Response(JSON.stringify(result.rows[0]), {
        headers: { 'Content-Type': 'application/json' },
      });
    } else {
      return new Response(
        JSON.stringify({ error: 'Product not found or no changes made' }),
        {
          status: 404,
          headers: { 'Content-Type': 'application/json' },
        }
      );
    }
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
  const id = params.id;

  try {
    const query = await client.query(
      'DELETE FROM Product WHERE id = $1 RETURNING *;',
      [id]
    );

    if (query.rowCount === 0) {
      return new Response(
        JSON.stringify({ message: 'No Product found with that ID.' }),
        {
          status: 404,
          headers: { 'Content-Type': 'application/json' },
        }
      );
    }

    return new Response(
      JSON.stringify({ message: 'Product deleted successfully.' }),
      {
        status: 200,
        headers: { 'Content-Type': 'application/json' },
      }
    );
  } catch (error) {
    console.error(error);
    return new Response(
      JSON.stringify({ error: 'Error performing the delete operation' }),
      {
        status: 500,
        headers: { 'Content-Type': 'application/json' },
      }
    );
  } finally {
    client.release();
  }
};
