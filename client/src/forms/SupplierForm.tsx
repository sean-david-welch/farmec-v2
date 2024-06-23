import utils from '../styles/Utils.module.css';
import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Supplier } from '../types/supplierTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { getFormFields } from '../utils/supplierFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	supplier?: Supplier;
}

const SupplierForm: React.FC<Props> = ({ id, supplier }) => {
	const [showForm, setShowForm] = useState(false);
	const formFields = supplier ? getFormFields(supplier) : getFormFields();

	const {
		mutateAsync: createSupplier,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Supplier>('suppliers');

	const {
		mutateAsync: updateSupplier,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Supplier>('suppliers', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitSupplier = id ? updateSupplier : createSupplier;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		const formData = new FormData(event.currentTarget as HTMLFormElement);

		const logoFile = formData.get('logo_image') as File;
		const marketingFile = formData.get('marketing_image') as File;

		const body: Supplier = {
			name: formData.get('name') as string,
			description: formData.get('description') as string,
			logo_image: logoFile ? logoFile.name : '',
			marketing_image: marketingFile ? marketingFile.name : '',
			social_facebook: formData.get('social_facebook') as string,
			social_twitter: formData.get('social_twitter') as string,
			social_instagram: formData.get('social_instagram') as string,
			social_youtube: formData.get('social_youtube') as string,
			social_linkedin: formData.get('social_linkedin') as string,
			social_website: formData.get('social_website') as string,
		};

		try {
			const response = await submitSupplier(body);

			if (logoFile) {
				const logoImageData = {
					imageFile: logoFile,
					presignedUrl: response.presignedLogoUrl,
				};
				await uploadFileToS3(logoImageData);
			}
			if (marketingFile) {
				const marketingImageData = {
					imageFile: marketingFile,
					presignedUrl: response.presignedMarketingUrl,
				};
				await uploadFileToS3(marketingImageData);
			}

			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('Error creating supplier:', error);
		}
	}

	if (createPending || updatingPending) return <Loading />;
	return (
		<section id="form">
			<button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
				{id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create Supplier'}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Supplier Form</h1>
					{formFields.map(field => (
						<div key={field.name}>
							<label htmlFor={field.name}>{field.label}</label>
							<input
								type={field.type}
								name={field.name}
								id={field.name}
								placeholder={field.placeholder}
								defaultValue={field.defaultValue}
							/>
						</div>
					))}
					<button className={utils.btnForm} type="submit">
						Submit
					</button>
				</form>

				{isError && <p>Error: {error?.message}</p>}
			</FormDialog>
		</section>
	);
};

export default SupplierForm;
