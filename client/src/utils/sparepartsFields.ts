import { Sparepart, Supplier } from '../types/supplierTypes';

export const getFormFields = (suppliers: Supplier[], sparepart?: Sparepart, fileLink?: boolean) => {
	const supplierOptions = Array.isArray(suppliers)
		? suppliers.map(supplier => ({
				label: supplier.name,
				value: supplier.id,
		  }))
		: [];

	return [
		{
			name: 'supplier_id',
			label: 'Supplier',
			type: 'select',
			options: supplierOptions,
			placeholder: 'Select supplier',
			defaultValue: sparepart?.supplier_id,
		},
		{
			name: 'name',
			label: 'Name',
			type: 'text',
			placeholder: 'Enter name',
			defaultValue: sparepart?.name,
		},
		{
			name: 'parts_image',
			label: 'Parts Image ',
			type: 'file',
			placeholder: 'Upload parts image',
		},
		{
			name: 'spare_parts_link_type',
			label: 'Spare Parts Link Type',
			type: 'radio',
			options: [
				{ value: 'url', label: 'URL' },
				{ value: 'file', label: 'File Upload' },
			],
			defaultValue: fileLink ? 'file' : 'url',
		},
		fileLink
			? {
					name: 'spare_parts_file_link',
					label: 'Spare Parts Link File',
					type: 'file',
					accept: '.pdf,.doc,.docx',
			  }
			: {
					name: 'spare_parts_url_link',
					label: 'Spare Parts Link',
					type: 'text',
					placeholder: 'Enter URL or select file',
					defaultValue: sparepart?.spare_parts_link || '',
			  },
	];
};
