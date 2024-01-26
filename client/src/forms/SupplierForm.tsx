import utils from '../styles/Utils.module.css';
import FormDialog from './FormDialog';

import { useState } from 'react';
import { Supplier } from '../types/supplierTypes';
import { useCreateSupplier, useUpdateSupplier } from '../hooks/supplierHooks';
import { getFormFields } from '../utils/supplierFormFields';
import { uploadFileToS3 } from '../utils/aws';

const SupplierForm: React.FC<{ id?: string }> = ({ id }) => {
    const formFields = getFormFields();
    const [showForm, setShowForm] = useState(false);

    const {
        mutateAsync: updateSupplier,
        isError: isUpdateError,
        error: updateError,
    } = id ? useUpdateSupplier(id) : { mutateAsync: () => {}, isError: false, error: null };

    const { mutateAsync: createSupplier, isError: isCreateError, error: createError } = useCreateSupplier();

    const isError = id ? isUpdateError : isCreateError;
    const error = id ? updateError : createError;

    const submitSupplier = id ? updateSupplier : createSupplier;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);

        const logoFile = formData.get('logo_image') as File;
        const marketingFile = formData.get('marketing_image') as File;

        const body: Supplier = {
            name: formData.get('name') as string,
            description: formData.get('description') as string,
            logo_image: logoFile ? logoFile.name : 'null',
            marketing_image: marketingFile ? marketingFile.name : 'null',
            social_facebook: formData.get('social_facebook') as string,
            social_twitter: formData.get('social_twitter') as string,
            social_instagram: formData.get('social_instagram') as string,
            social_youtube: formData.get('social_youtube') as string,
            social_linkedin: formData.get('social_linkedin') as string,
            social_website: formData.get('social_website') as string,
        };

        try {
            const response = await submitSupplier(body);

            console.log(response);

            if (logoFile) {
                const logoImageData = {
                    imageFile: logoFile,
                    presignedUrl: response.presignedLogoUrl,
                };
                await uploadFileToS3(logoImageData);
            }
            if (marketingFile) {
                const marketingImageData = {
                    imageFile: marketingFile,
                    presignedUrl: response.presignedMarketingUrl,
                };
                await uploadFileToS3(marketingImageData);
            }
        } catch (error) {
            console.error('Error creating supplier:', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <img src="/icons/edit.svg" alt="edit button" /> : 'Create Supplier'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Supplier Form</h1>
                    {formFields.map(field => (
                        <div key={field.name}>
                            <label htmlFor={field.name}>{field.label}</label>
                            <input
                                type={field.type}
                                name={field.name}
                                id={field.name}
                                placeholder={field.placeholder}
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

export default SupplierForm;
