import { Pool } from 'pg';

import secrets from './secrets';

const connectionString = secrets.database_url;

export const pool = new Pool({
  connectionString: connectionString,
});
