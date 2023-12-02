import { Pool } from 'pg';

const connectionString = import.meta.env.DATABASE_URL;

export const pool = new Pool({
  connectionString: connectionString,
});
