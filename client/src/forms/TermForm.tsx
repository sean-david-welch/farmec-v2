import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Terms } from '../types/aboutTypes';

import { useMutateResource } from '../hooks/genericHooks';
import { termsFormFields } from '../utils/aboutFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	term?: Terms;
}

const TermForm: React.FC<Props> = ({ id, term }) => {
	const [showForm, setShowForm] = useState(false);
	const formFields = term ? termsFormFields(term) : termsFormFields();

	const {
		mutateAsync: createTerm,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Terms>('terms');

	const {
		mutateAsync: updateTerm,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Terms>('terms', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitTerm = id ? updateTerm : createTerm;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);

		const body: Terms = {
			title: formData.get('title') as string,
			body: formData.get('body') as string,
		};

		try {
			const response = await submitTerm(body);
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('error creating term', error);
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
						Create Term
						<FontAwesomeIcon icon={faPenToSquare} />
					</div>
				)}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Terms Form</h1>
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

export default TermForm;
