import type { APIRoute } from 'astro';
import { app } from '../../../lib/firebase-server';
import { getAuth } from 'firebase-admin/auth';

import { pool } from '../../../database/connection';

export const GET: APIRoute = async ({ request, cookies, redirect }) => {
  const auth = getAuth(app);
  const client = await pool.connect();

  const idToken = request.headers.get('Authorization')?.split('Bearer ')[1];
  if (!idToken) {
    return new Response('No token found', { status: 401 });
  }

  try {
    const decodedToken = await auth.verifyIdToken(idToken);
    const uid = decodedToken.uid;

    const query = await client.query('SELECT * FROM "users" WHERE "uid" = $1;', [uid]);
    const user = query.rows[0];
    console.log('User:', user);

    const threeDays = 60 * 60 * 24 * 3 * 1000;
    const sessionCookie = await auth.createSessionCookie(idToken, {
      expiresIn: threeDays,
    });

    cookies.set('session', sessionCookie, {
      path: '/',
      httpOnly: true,
    });

    return new Response(JSON.stringify(user), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    return new Response('Invalid token', { status: 401 });
  } finally {
    client.release();
  }
};
