import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Timeline } from '../types/aboutTypes';

import { useMutateResource } from '../hooks/genericHooks';
import { timelineFormFields } from '../utils/aboutFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	timeline?: Timeline;
}

const TimelineForm: React.FC<Props> = ({ id, timeline }) => {
	const [showForm, setShowForm] = useState(false);
	const formFields = timeline ? timelineFormFields(timeline) : timelineFormFields();

	const {
		mutateAsync: createTimeline,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Timeline>('timelines');

	const {
		mutateAsync: updateTimeline,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Timeline>('timelines', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitTimeline = id ? updateTimeline : createTimeline;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);

		const body: Timeline = {
			title: formData.get('title') as string,
			date: formData.get('date') as string,
			body: formData.get('body') as string,
		};

		try {
			const response = await submitTimeline(body);
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('error creating timeline', error);
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
						Create Timeline Event
						<FontAwesomeIcon icon={faPenToSquare} />
					</div>
				)}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Timeline Form</h1>
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

export default TimelineForm;
