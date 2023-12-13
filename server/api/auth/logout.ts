import type { APIRoute } from 'astro';

export const GET: APIRoute = async ({ cookies }) => {
  cookies.delete('session', {
    path: '/',
  });

  return new Response(JSON.stringify({ Message: 'User Logged Out' }), {
    status: 200,
    headers: {
      'Content-Type': 'application/json',
    },
  });
};
