import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { Blog } from '../types/blogTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { blogFormFields } from '../utils/blogFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
	id?: string;
	blog?: Blog;
}

const BlogForm: React.FC<Props> = ({ id, blog }) => {
	const [showForm, setShowForm] = useState(false);
	const formFields = blog ? blogFormFields(blog) : blogFormFields();

	const {
		mutateAsync: createBlog,
		isError: isCreateError,
		error: createError,
		isPending: createPending,
	} = useMutateResource<Blog>('blogs');

	const {
		mutateAsync: updateBlog,
		isError: isUpdateError,
		error: updateError,
		isPending: updatingPending,
	} = useMutateResource<Blog>('blogs', id);

	const error = id ? updateError : createError;
	const isError = id ? isUpdateError : isCreateError;
	const submitBlog = id ? updateBlog : createBlog;

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();

		const formData = new FormData(event.currentTarget as HTMLFormElement);
		const imageFile = formData.get('main_image') as File;

		const body: Blog = {
			title: formData.get('title') as string,
			date: formData.get('date') as string,
			main_image: imageFile ? imageFile.name : 'null',
			subheading: formData.get('subheading') as string,
			body: formData.get('body') as string,
		};

		try {
			const response = await submitBlog(body);

			if (imageFile) {
				const imageData = {
					imageFile: imageFile,
					presignedUrl: response.presignedUrl,
				};
				await uploadFileToS3(imageData);
			}
			response && !isError && setShowForm(false);
		} catch (error) {
			console.error('error creating blog', error);
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
						Create Blog
						<FontAwesomeIcon icon={faPenToSquare} />
					</div>
				)}
			</button>

			<FormDialog visible={showForm} onClose={() => setShowForm(false)}>
				<form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
					<h1 className={utils.mainHeading}>Blog Form</h1>
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

export default BlogForm;
