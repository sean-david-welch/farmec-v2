import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Privacy } from '../types/aboutTypes';

import { useMutateResource } from '../hooks/genericHooks';
import { privacyFormFields } from '../utils/aboutFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	privacy?: Privacy;
}

const PrivacyForm: React.FC<Props> = ({ id, privacy }) => {
	const [showForm, setShowForm] = useState(false);
	const formFields = privacy ? privacyFormFields(privacy) : privacyFormFields();

	const {
		mutateAsync: createPrivacy,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Privacy>('privacys');

	const {
		mutateAsync: updatePrivacy,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Privacy>('privacys', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitPrivacy = id ? updatePrivacy : createPrivacy;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);

		const body: Privacy = {
			title: formData.get('title') as string,
			body: formData.get('body') as string,
		};

		try {
			const response = await submitPrivacy(body);
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('error creating privacy', error);
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
						Create Privacy
						<FontAwesomeIcon icon={faPenToSquare} />
					</div>
				)}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Privacy Form</h1>
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

export default PrivacyForm;
