import type { MiddlewareHandler, MiddlewareNext, APIContext } from 'astro';
import { verifyToken } from '../utils/admin';

const middleware: MiddlewareHandler<Promise<void>> = async (
  context: APIContext,
  next: MiddlewareNext<Promise<void>>
) => {
  console.log('middle running');

  if (['POST', 'PUT', 'DELETE'].includes(context.request.method)) {
    console.log('checked method');

    await verifyToken(context.request);
  }

  return next();
};

export const onRequest = middleware;
