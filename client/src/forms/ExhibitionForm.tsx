import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';

import { useState } from 'react';
import { Exhibition } from '../types/blogTypes';

import { useMutateResource } from '../hooks/genericHooks';
import { exhibitionFormFields } from '../utils/blogFields';

interface Props {
    id?: string;
    exhibition?: Exhibition;
}

const ExhibitionForm: React.FC<Props> = ({ id, exhibition }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = exhibition ? exhibitionFormFields(exhibition) : exhibitionFormFields();

    const {
        mutateAsync: createExhibition,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<Exhibition>('exhibitions');

    const {
        mutateAsync: updateExhibition,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<Exhibition>('exhibitions', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitExhibition = id ? updateExhibition : createExhibition;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);

        const body: Exhibition = {
            title: formData.get('title') as string,
            date: formData.get('date') as string,
            location: formData.get('location') as string,
            info: formData.get('info') as string,
        };

        try {
            const response = await submitExhibition(body);
            response.ok ? setShowForm(false) : console.error('failed with response:', response);
        } catch (error) {
            console.error('error creating exhibition', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <img src="/icons/edit.svg" alt="edit button" /> : 'Create Exhibition'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Exhibition Form</h1>
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

export default ExhibitionForm;
