import utils from '../styles/Utils.module.css';
import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { useMutateResource } from '../hooks/genericHooks';
import { uploadFileToS3 } from '../lib/aws';
import { Machine, Supplier } from '../types/supplierTypes';
import { getFormFields } from '../utils/machineFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	machine?: Machine;
	suppliers: Supplier[];
}

const MachineFrom: React.FC<Props> = ({ id, machine, suppliers }) => {
	const [showForm, setShowForm] = useState(false);

	const formFields = machine ? getFormFields(suppliers, machine) : getFormFields(suppliers);

	const {
		mutateAsync: createMachine,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Machine>('machines');

	const {
		mutateAsync: updateMachine,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Machine>('machines', id);

	const isError = id ? isUpdateError : isCreateError;
	const error = id ? updateError : createError;

	const submitMachine = id ? updateMachine : createMachine;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);

		const imageFile = formData.get('machine_image') as File;

		const body: Machine = {
			supplier_id: formData.get('supplier_id') as string,
			name: formData.get('name') as string,
			machine_image: imageFile ? imageFile.name : 'null',
			description: formData.get('description') as string,
			machine_link: formData.get('machine_link') as string,
		};

		try {
			const response = await submitMachine(body);

			if (imageFile) {
				const machineImageData = {
					imageFile: imageFile,
					presignedUrl: response.presignedUrl,
				};
				await uploadFileToS3(machineImageData);
			}
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('Error creating machine', error);
		}
	}

	if (createPending || updatingPending) return <Loading />;
	return (
		<section id="form">
			<button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
				{id ? (
					<FontAwesomeIcon icon={faPenToSquare} />
				) : (
					<div>
						Create Machine
						<FontAwesomeIcon icon={faPenToSquare} />
					</div>
				)}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Machine Form</h1>
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

export default MachineFrom;
