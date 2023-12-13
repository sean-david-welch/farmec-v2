import fastify from 'fastify';
import fastifyCors from '@fastify/cors';
import fastifyCompress from '@fastify/compress';
import fastifyCookie from '@fastify/cookie';
import fastifyPostgres from '@fastify/postgres';
import fastifyFormbody from '@fastify/formbody';

const app = fastify();

app.register(fastifyCors, {
  origin: 'https://bounce-frontend-vite.onrender.com',
  credentials: true,
});
app.register(fastifyCompress);
app.register(fastifyCookie);
app.register(fastifyFormbody);

// app.register(router, { prefix: '/api' });

app.get('/docs.json', (request, reply) => {
  reply.header('Content-Type', 'application/json');
});

app
  .listen({
    port: 8080,
    host: '0.0.0.0',
  })
  .then(address => {
    console.log(`Server running on ${address}`);
  })
  .catch(err => {
    console.error(err);
    process.exit(1);
  });
