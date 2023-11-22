import { getAuth } from 'firebase-admin/auth';
import { app } from '../lib/firebase-server';

export const verifyToken = async (request: Request) => {
  const cookiesHeader = request.headers.get('cookie') || '';

  const cookies = new Map(
    cookiesHeader.split('; ').map(cookie => {
      const [name, ...value] = cookie.split('=');
      return [name, value.join('=')];
    })
  );

  const token = cookies.get('session');
  console.log('Token: ', token);

  if (!token) {
    return new Response(JSON.stringify({ error: 'No token provided' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' },
    });
  }

  try {
    const auth = getAuth(app);

    const decodedToken = await auth.verifySessionCookie(token);

    return decodedToken;
  } catch (error) {
    console.error('Error verifying token:', error);
    throw error;
  }
};
