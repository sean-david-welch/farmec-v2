import Supplier from 'types/supplier';
import SupplierService from 'services/suppliers';

import { generatePresignedUrl, generateDeletePresignedUrl } from '../utils/aws';
import { FastifyReply, FastifyRequest } from 'fastify';

class SuppliersController {
  private folder: string = 'suppliers';
  private supplierService: SupplierService;

  constructor(supplierService: SupplierService) {
    this.supplierService = supplierService;
  }

  async getSuppliers(request: FastifyRequest, reply: FastifyReply) {
    try {
      const suppliers = await this.supplierService.getSuppliers();

      return reply.code(200).send(suppliers);
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async createSupplier(request: FastifyRequest<{ Body: Supplier }>, reply: FastifyReply) {
    const data = request.body;
    try {
      const { logoUploadUrl, logoImageUrl } = await generatePresignedUrl(this.folder, data.logo_image);
      const { marketingUploadUrl, marketingImageUrl } = await generatePresignedUrl(this.folder, data.marketing_image);

      const supplierData = {
        ...data,
        logoImageUrl,
        marketingImageUrl,
      };

      const supplier = await this.supplierService.createSupplier(supplierData);

      return reply.code(201).send({
        supplier,
        logoUploadUrl: logoUploadUrl,
        marketingUploadUrl: marketingUploadUrl,
      });
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async updateSupplier(request: FastifyRequest<{ Params: { id: string }; Body: Supplier }>, reply: FastifyReply) {
    const id = request.params.id;
    const data = request.body;

    const images = {
      logo: data.logo_image as string,
      marketing: data.logo_image as string,
    };

    try {
      const { logoUploadUrl, logoImageUrl } = await generatePresignedUrl(this.folder, data.logo_image);
      const { marketingUploadUrl, marketingImageUrl } = await generatePresignedUrl(this.folder, data.marketing_image);

      const supplierData = {
        ...data,
        logoImageUrl,
        marketingImageUrl,
      };

      const supplier = await this.supplierService.updateSupplier(id, supplierData);

      return reply.code(200).send({
        supplier,
        logoUploadUrl: logoUploadUrl,
        marketingUploadUrl: marketingUploadUrl,
      });
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }

  async deleteSupplier(request: FastifyRequest<{ Params: { id: string } }>, reply: FastifyReply) {
    const id = request.params.id;

    try {
      const supplier = await this.supplierService.deleteSupplier(id);

      const images = {
        logo: supplier.logo_image as string,
        marketing: supplier.marketing_image as string,
      };

      const deleteLogoUrl = await generateDeletePresignedUrl(images.logo);
      const deleteMarketingUrl = await generateDeletePresignedUrl(images.marketing);

      return reply.code(200).send({
        message: 'Supplier deleted',
        supplier,
        deleteLogoUrl,
        deleteMarketingUrl,
      });
    } catch (error) {
      console.error('Controller Error:', error);

      return reply.code(500).send({ error: 'Internal Server Error' });
    }
  }
}

export default SuppliersController;
