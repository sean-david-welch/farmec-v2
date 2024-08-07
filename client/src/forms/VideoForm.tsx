import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Supplier } from '../types/supplierTypes';

import { useMutateResource } from '../hooks/genericHooks';
import { Video, VideoWebUrl } from '../types/videoTypes';
import { getFormFields } from '../utils/videoFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	video?: Video;
	suppliers: Supplier[];
}

const VideoForm: React.FC<Props> = ({ id, video, suppliers }) => {
	const [showForm, setShowForm] = useState(false);
	const formFields = video ? getFormFields(suppliers, video) : getFormFields(suppliers);

	const {
		mutateAsync: createVideo,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<VideoWebUrl>('videos');

	const {
		mutateAsync: updateVideo,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<VideoWebUrl>('videos', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitVideo = id ? updateVideo : createVideo;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);

		const body: VideoWebUrl = {
			supplier_id: formData.get('supplier_id') as string,
			web_url: formData.get('web_url') as string,
		};

		try {
			const response = await submitVideo(body);
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('error creating video', error);
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
						Create Video
						<FontAwesomeIcon icon={faPenToSquare} />
					</div>
				)}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Videos Form</h1>

					{formFields.map(field => (
						<div key={field.name}>
							<label htmlFor={field.name}>{field.label}</label>
							{field.type === 'select' ? (
								<select name={field.name} id={field.name}>
									{field.options?.map(option => (
										<option key={option.value} value={option.value}>
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

export default VideoForm;
