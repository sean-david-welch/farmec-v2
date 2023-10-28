import { pool } from '../../database/connection';

export const GET = async (): Promise<Response> => {
  const client = await pool.connect();

  try {
    const query = await client.query('select * from "Supplier"');
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

export const POST = async (request: Request): Promise<Response> => {
  const client = await pool.connect();

  try {
    const data = await request.json();

    const sql = `INSERT INTO Supplier(id, name, logo_image, marketing_image, description, social_facebook, social_twitter, social_instagram, social_youtube, social_linkedin, social_website, created)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    RETURNING *;
    `;

    const values = [
      data.id,
      data.name,
      data.logo_image,
      data.marketing_image,
      data.description,
      data.social_facebook,
      data.social_twitter,
      data.social_instagram,
      data.social_youtube,
      data.social_linkedin,
      data.social_website,
      data.created,
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
