import type { APIRoute } from 'astro';

import { app } from '../../../lib/firebase-server';
import { getAuth } from 'firebase-admin/auth';

export const GET: APIRoute = async ({ request, cookies }) => {
  const auth = getAuth(app);

  const idToken = request.headers.get('Authorization')?.split('Bearer ')[1];
  if (!idToken) {
    return new Response('No token found', { status: 401 });
  }

  try {
    const decodedToken = await auth.verifyIdToken(idToken);
    const user = {
      uid: decodedToken.uid,
      email: decodedToken.email,
    };

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
  }
};
