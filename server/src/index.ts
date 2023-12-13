import fastify from 'fastify';
import fastifyCors from '@fastify/cors';
import fastifyCompress from '@fastify/compress';
import fastifyCookie from '@fastify/cookie';
import fastifyPostgres from '@fastify/postgres';
import fastifyFormbody from '@fastify/formbody';

import secrets from './utils/secrets';
import mainRouter from './router/router';

const app = fastify();

app.register(fastifyCors, {
  origin: 'http://localhost:3000',
  credentials: true,
});

app.register(fastifyCompress);
app.register(fastifyCookie);
app.register(fastifyFormbody);

app.register(fastifyPostgres, {
  connectionString: secrets.database_url,
});

app.register(mainRouter, { prefix: '/api' });

app.get('/', (request, reply) => {
  reply.header('Content-Type', 'application/json');
  reply.send('Fastify App Running');
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
