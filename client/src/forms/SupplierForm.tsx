import utils from '../styles/Utils.module.css';
import FormDialog from './FormDialog';

import { useState } from 'react';
import { Supplier } from '../types/supplierTypes';
import { useCreateSupplier } from '../hooks/supplierHooks';
import { getFormFields } from '../utils/supplierFormFields';

const SupplierForm = () => {
    const formFields = getFormFields();

    const [showForm, setShowForm] = useState(false);

    const { mutateAsync: createSupplier, isError, error } = useCreateSupplier();

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
            await createSupplier(body);
        } catch (error) {
            console.error('Error creating supplier:', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                Create Supplier
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

                {isError && <p>Error: {error.message}</p>}
            </FormDialog>
        </section>
    );
};

export default SupplierForm;
