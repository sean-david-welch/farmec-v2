import type { APIRoute } from 'astro';
import { app } from '../../../lib/firebase-server';
import { getAuth } from 'firebase-admin/auth';

export const GET: APIRoute = async ({ request, cookies, redirect }) => {
  const auth = getAuth(app);

  /* Get token from request headers */
  const idToken = request.headers.get('Authorization')?.split('Bearer ')[1];
  if (!idToken) {
    return new Response('No token found', { status: 401 });
  }

  /* Verify id token */
  try {
    await auth.verifyIdToken(idToken);
  } catch (error) {
    return new Response('Invalid token', { status: 401 });
  }

  /* Create and set session cookie */
  const threeDays = 60 * 60 * 24 * 3 * 1000;
  const sessionCookie = await auth.createSessionCookie(idToken, {
    expiresIn: threeDays,
  });

  cookies.set('session', sessionCookie, {
    path: '/',
    httpOnly: true,
  });

  return new Response(JSON.stringify({ success: true }), {
    status: 200,
    headers: { 'Content-Type': 'application/json' },
  });
};
