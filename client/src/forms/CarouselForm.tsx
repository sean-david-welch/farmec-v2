import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';

import { useState } from 'react';
import { Carousel } from '../types/miscTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { getFormFields } from '../utils/carouselFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
    id?: string;
    carousel?: Carousel;
}

const CarouselForm: React.FC<Props> = ({ id, carousel }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = carousel ? getFormFields(carousel) : getFormFields();

    const {
        mutateAsync: createCarousel,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<Carousel>('carousels');

    const {
        mutateAsync: updateCarousel,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<Carousel>('carousels', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitCarousel = id ? updateCarousel : createCarousel;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);
        const imageFile = formData.get('image') as File;

        const body: Carousel = {
            name: formData.get('name') as string,
            image: imageFile ? imageFile.name : 'null',
        };

        try {
            const response = await submitCarousel(body);

            if (imageFile) {
                const imageData = {
                    imageFile: imageFile,
                    presignedUrl: response.presignedUrl,
                };
                await uploadFileToS3(imageData);
            }

            response && !isError && setShowForm(false);
        } catch (error) {
            console.error('error creating carousel', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create Carousel'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Carousel Form</h1>
                    {formFields.map((field) => (
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

export default CarouselForm;
