import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Sparepart, Supplier } from '../types/supplierTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { getFormFields } from '../utils/sparepartsFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	sparepart?: Sparepart;
	suppliers: Supplier[];
}

const SparepartForm: React.FC<Props> = ({ id, sparepart, suppliers }) => {
	const [showForm, setShowForm] = useState(false);
	const [fileLink, setFileLink] = useState(false);
	const formFields = sparepart
		? getFormFields(suppliers, sparepart, fileLink)
		: getFormFields(suppliers, undefined, fileLink);

	const {
		mutateAsync: createSparepart,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Sparepart>('spareparts');

	const {
		mutateAsync: updateSparepart,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Sparepart>('spareparts', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitSparepart = id ? updateSparepart : createSparepart;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);
		const imageFile = formData.get('parts_image') as File;

		const body: Sparepart = {
			supplier_id: formData.get('supplier_id') as string,
			name: formData.get('name') as string,
			parts_image: imageFile ? imageFile.name : 'null',
			spare_parts_link: formData.get('spare_parts_link') as string,
		};

		try {
			const response = await submitSparepart(body);

			if (imageFile) {
				const imageData = {
					imageFile: imageFile,
					presignedUrl: response.presignedUrl,
				};
				await uploadFileToS3(imageData);
			}
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('error creating sparepart', error);
		}
	}

	if (createPending || updatingPending) return <Loading />;
	return (
		<section id="form">
			<button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
				{id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create Sparepart'}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Sparepart Form</h1>
					{formFields.map(field => (
						<div key={field.name}>
							<label htmlFor={field.name}>{field.label}</label>
							{field.type === 'select' ? (
								<select name={field.name} id={field.name}>
									{field.options?.map(option => (
										<option
											key={option.value}
											value={option.value}
											defaultValue={field.defaultValue}>
											{option.label}
										</option>
									))}
								</select>
							) : field.type === 'radio' ? (
								<div onChange={e => setFileLink((e.target as HTMLInputElement).value === 'file')}>
									{field.options?.map(option => (
										<label key={option.value}>
											<input type="radio" name={field.name} value={option.value} />
											{option.label}
										</label>
									))}
								</div>
							) : (
								<input
									type={field.type}
									name={field.name}
									id={field.name}
									placeholder={field.placeholder}
									defaultValue={field.defaultValue}
								/>
							)}
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

export default SparepartForm;
