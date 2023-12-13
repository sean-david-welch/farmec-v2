import { v4 as uuidv4 } from 'uuid';
import type { APIRoute } from 'astro';
import { pool } from '../../../database/connection';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query('select * from "Registration"');
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

    const sql = `
    INSERT INTO Registration (
        id, 
        dealer_name, 
        dealer_address, 
        owner_name, 
        owner_address, 
        machine_model, 
        serial_number, 
        install_date, 
        invoice_number, 
        complete_supply, 
        pdi_complete, 
        pto_correct, 
        machine_test_run, 
        safety_induction, 
        operator_handbook, 
        date, 
        completed_by, 
        created
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
    )
    RETURNING *;
    `;

    const uuid = uuidv4();
    const currentDateTime = new Date().toISOString();

    const values = [
      uuid,
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
      currentDateTime,
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
