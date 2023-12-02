import type { MiddlewareNext, APIContext } from 'astro';
import { verifyToken } from '../utils/admin';

type CustomMiddlewareHandler = (
  context: APIContext,
  next: MiddlewareNext<Promise<void>>
) => Promise<void | Response> | void | Response;

const middleware: CustomMiddlewareHandler = async (context: APIContext, next: MiddlewareNext<Promise<void>>) => {
  if (['POST', 'PUT', 'DELETE'].includes(context.request.method)) {
    const result = await verifyToken(context.request);

    if (result instanceof Response) {
      return result;
    }

    const url = new URL(context.request.url);
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

  return next();
};

export const onRequest = middleware;
