const GetSuppliers = async (request, reply) => {
  try {
    const suppliers = await supplierService.getSuppliers();
    return reply.code(200).send(suppliers);
  } catch (error) {
    console.error(error);
    return reply.code(500).send({ error: 'Internal Server Error' });
  }
};
