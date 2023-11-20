import { getAuth } from 'firebase-admin/auth';
import { app } from '../lib/firebase-server';

export const verifyToken = async (request: Request) => {
  const cookiesHeader = request.headers.get('cookie') || '';
  console.log(cookiesHeader);

  const cookies = new Map(
    cookiesHeader.split('; ').map(cookie => {
      const [name, ...value] = cookie.split('=');
      return [name, value.join('=')];
    })
  );

  console.log(cookies);

  const token = cookies.get('session');
  console.log(token);

  if (!token) {
    return new Response(JSON.stringify({ error: 'No token provided' }), {
      status: 401,
      headers: { 'Content-Type': 'application/json' },
    });
  }

  try {
    const auth = getAuth(app);

    const decodedToken = await auth.verifyIdToken(token);

    return decodedToken;
  } catch (error) {
    throw new Error('Invalid token');
  }
};
