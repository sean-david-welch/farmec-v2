const SuppliersService = async () => {
  const client = await pool.connect();
  try {
    const query = await client.query('select * from "Supplier"');
    return query.rows;
  } catch (error) {
    console.error(error);
    throw new Error('Database query failed');
  } finally {
    client.release();
  }
};
