import { FastifyRequest, FastifyReply } from 'fastify';
import { verifyToken } from '../utils/firebase';

export const isAuthenticated = async (request: FastifyRequest, reply: FastifyReply, done: Function) => {
  try {
    if (['POST', 'PUT', 'DELETE'].includes(request.method)) {
      const result = await verifyToken(request, reply);

      const url = new URL(request.url);
      const path = url.pathname;

      const isAdmin = result.admin === true;

      if (!isAdmin && !path.startsWith('/api/services/')) {
        reply.code(403).send({ error: 'Access denied' });
        return;
      }
    }

    done();
  } catch (error) {
    console.log(error);
    reply.code(400).send();
    return;
  }
};

export default isAuthenticated;
