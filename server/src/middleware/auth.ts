import { FastifyRequest, FastifyReply } from 'fastify';
import { verifyToken } from '../utils/firebase';

export const isAuthenticated = async (request: FastifyRequest, reply: FastifyReply, done: Function) => {
  try {
    if (['POST', 'PUT', 'DELETE'].includes(request.method)) {
      const result = await verifyToken(request, reply);

      if (result instanceof Response) {
        return result;
      }

      const url = new URL(request.url);
      const path = url.pathname;
      console.log('path', path);

      const isAdmin = result.admin === true;

      if (!isAdmin && !path.startsWith('/api/services/')) {
        return new Response(JSON.stringify({ error: 'Access denied' }), {
          status: 403,
          headers: { 'Content-Type': 'application/json' },
        });
      }
    }

    done();
  } catch (error) {
    console.log(error);
    reply.code(400).send();
  }
};

export default isAuthenticated;
