import { app } from '../lib/firebase';
import { getAuth } from 'firebase-admin/auth';
import { FastifyRequest, FastifyReply } from 'fastify';

export const verifyToken = async (request: FastifyRequest, reply: FastifyReply) => {
  const cookies = request.cookies;

  const token = cookies.session;
  console.log('token in middleware', token);

  if (!token) {
    reply.code(401).send({ error: 'No token provided' });
    return;
  }

  try {
    const auth = getAuth(app);

    const decodedToken = await auth.verifySessionCookie(token);
    const isAdmin = decodedToken.admin === true;

    console.log('Admin:', isAdmin);
    console.log('token:', decodedToken);

    return decodedToken;
  } catch (error) {
    console.error('Error verifying token:', error);
    throw error;
  }
};
