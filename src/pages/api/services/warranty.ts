import { v4 as uuidv4 } from 'uuid';
import type { APIRoute } from 'astro';
import { pool } from '../../../database/connection';

export const GET: APIRoute = async ({ params }): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query(`
    SELECT wc.*, pr.* 
    FROM "WarrantyClaim" wc
    LEFT JOIN 
    "PartsRequired" pr ON wc.id = pr."warrantyId";
`);
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
    await client.query('BEGIN');
    const data = await request.json();

    const warrantyClaimID = uuidv4();
    const currentDateTime = new Date().toISOString();

    const warrantyClaimSQL = `
    INSERT INTO WarrantyClaim (
        id, 
        dealer, 
        dealer_contact, 
        owner_name, 
        machine_model, 
        serial_number, 
        install_date, 
        failure_date,
        repair_date,
        failure_details,
        repair_details,
        labour_hours,
        completed_by,
        created
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
    )
    RETURNING *;
  `;
    const warrantyClaimValues = [
      warrantyClaimID,
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
      currentDateTime,
    ];
    const warrantyClaimResult = await client.query(
      warrantyClaimSQL,
      warrantyClaimValues
    );

    if (data.partsRequired && data.partsRequired.length > 0) {
      for (const part of data.partsRequired) {
        const partsRequiredID = uuidv4();

        const partsRequiredSQL = `
        INSERT INTO PartsRequired (
            id,
            warrantyId,
            part_number,
            quantity_needed,
            invoice_number,
            description
        ) VALUES (
            $1, $2, $3, $4, $5, $6
        );
        RETURNING *`;

        const partsRequiredValues = [
          partsRequiredID,
          warrantyClaimID,
          part.part_number,
          part.quantity_needed,
          part.invoice_number,
          part.description,
        ];
        await client.query(partsRequiredSQL, partsRequiredValues);
      }
    }

    await client.query('COMMIT');

    return new Response(JSON.stringify(warrantyClaimResult.rows[0]), {
      headers: { 'Content-Type': 'application/json' },
    });
  } catch (error) {
    console.error(error);
    await client.query('ROLLBACK');

    return new Response(JSON.stringify({ error: 'Database query failed' }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' },
    });
  } finally {
    client.release();
  }
};
