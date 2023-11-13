import type { APIRoute } from 'astro';
import { pool } from '../../../../database/connection';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const query = await client.query(
      `SELECT wc.*, pr.* 
      FROM "WarrantyClaim" wc
      LEFT JOIN "PartsRequired" pr ON wc.id = pr."warrantyId"
      WHERE wc.id = $1;`,
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

    const sql = `
          UPDATE WarrantyClaim SET 
          dealer = $1, 
          dealer_contact = $2, 
          owner_name = $3, 
          machine_model = $4, 
          serial_number = $5, 
          install_date = $6, 
          failure_date = $7,
          repair_date = $8,
          failure_details = $9,
          repair_details = $10,
          labour_hours = $11,
          completed_by = $12,
          WHERE id = $1
          RETURNING *;`;

    const values = [
      id,
      data.dealer,
      data.dealer_contact,
      data.owner_name,
      data.machine_model,
      data.serial_number,
      data.install_date,
      data.failure_date,
      data.repair_date,
      data.failure_details,
      data.repair_details,
      data.labour_hours,
      data.completed_by,
    ];

    const result = await client.query(sql, values);

    if (result.rows.length > 0) {
      return new Response(JSON.stringify(result.rows[0]), {
        headers: { 'Content-Type': 'application/json' },
      });
    } else {
      return new Response(
        JSON.stringify({ error: 'Registration not found or no changes made' }),
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
      'DELETE FROM Registration WHERE id = $1 RETURNING *;',
      [id]
    );

    if (query.rowCount === 0) {
      return new Response(
        JSON.stringify({ message: 'No Registration found with that ID.' }),
        {
          status: 404,
          headers: { 'Content-Type': 'application/json' },
        }
      );
    }

    return new Response(
      JSON.stringify({ message: 'Blog deleted successfully.' }),
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
