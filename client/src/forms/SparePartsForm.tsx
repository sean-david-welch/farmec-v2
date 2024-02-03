import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';

import { useState } from 'react';
import { Sparepart, Supplier } from '../types/supplierTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { getFormFields } from '../utils/sparepartsFields';

interface Props {
    id?: string;
    sparepart?: Sparepart;
    suppliers: Supplier[];
}

const SparepartForm: React.FC<Props> = ({ id, sparepart, suppliers }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = sparepart ? getFormFields(suppliers, sparepart) : getFormFields(suppliers);

    const {
        mutateAsync: createSparepart,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<Sparepart>('spareparts');

    const {
        mutateAsync: updateSparepart,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<Sparepart>('spareparts', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitSparepart = id ? updateSparepart : createSparepart;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);
        const imageFile = formData.get('parts_image') as File;

        const body: Sparepart = {
            supplierId: formData.get('supplierId') as string,
            name: formData.get('name') as string,
            parts_image: formData.get('parts_image') as string,
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
            response.ok ? setShowForm(false) : console.error('failed with response:', response);
        } catch (error) {
            console.error('error creating sparepart', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <img src="/icons/edit.svg" alt="edit button" /> : 'Create Sparepart'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Sparepart Form</h1>
                    {formFields.map((field) => (
                        <div key={field.name}>
                            <label htmlFor={field.name}>{field.label}</label>
                            {field.type === 'select' ? (
                                <select name={field.name} id={field.name}>
                                    {field.options?.map((option) => (
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

export default SparepartForm;
