import type { APIRoute } from 'astro';

import { getAuth } from 'firebase-admin/auth';
import { app } from '../../../lib/firebase-server';

export const GET: APIRoute = async ({ request }): Promise<Response> => {
  const cookiesHeader = request.headers.get('cookie') || '';

  const cookies = new Map(
    cookiesHeader.split('; ').map(cookie => {
      const [name, ...value] = cookie.split('=');
      return [name, value.join('=')];
    })
  );

  const token = cookies.get('session');
  console.log('token', token);

  if (!token) {
    return new Response(JSON.stringify({ userLoggedIn: false, isAdmin: false }), {
      headers: {
        'Content-Type': 'application/json',
      },
    });
  }

  try {
    const auth = getAuth(app);
    const decodedToken = await auth.verifySessionCookie(token);
    const isAdmin = decodedToken.admin === true;

    console.log('Admin api', isAdmin);
    console.log('token', decodedToken);

    return new Response(JSON.stringify({ userLoggedIn: true, isAdmin }), {
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    console.error('Error verifying token:', error);
    return new Response(JSON.stringify({ error: 'Error processing request', userLoggedIn: false, isAdmin: false }), {
      headers: { 'Content-Type': 'application/json' },
    });
  }
};
