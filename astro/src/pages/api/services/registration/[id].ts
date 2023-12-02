import type { APIRoute } from 'astro';
import { pool } from '../../../../database/connection';

export const GET: APIRoute = async ({ params }) => {
  const client = await pool.connect();
  const id = params.id;

  try {
    const query = await client.query(
      'SELECT * FROM "Registration" WHERE "id" = $1;',
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
        UPDATE Registration SET 
        dealer_name = $2, 
        dealer_address = $3, 
        owner_name = $4, 
        owner_address = $5, 
        machine_model = $6, 
        serial_number = $7, 
        install_date = $8, 
        invoice_number = $9, 
        complete_supply = $10, 
        pdi_complete = $11, 
        pto_correct = $12, 
        machine_test_run = $13, 
        safety_induction = $14, 
        operator_handbook = $15, 
        date = $16, 
        completed_by = $17, 
        WHERE id = $1
        RETURNING *;`;

    const values = [
      id,
      data.dealer_name,
      data.dealer_address,
      data.owner_name,
      data.owner_address,
      data.machine_model,
      data.serial_number,
      data.install_date,
      data.invoice_number,
      data.complete_supply,
      data.pdi_complete,
      data.pto_correct,
      data.machine_test_run,
      data.safety_induction,
      data.operator_handbook,
      data.date,
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
