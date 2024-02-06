import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';

import { useState } from 'react';
import { LineItem } from '../types/miscTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { getFormFields } from '../utils/lineItemFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
    id?: string;
    lineItem?: LineItem;
}

const LineItemForm: React.FC<Props> = ({ id, lineItem }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = lineItem ? getFormFields(lineItem) : getFormFields();

    const {
        mutateAsync: createLineItem,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<LineItem>('lineitems');

    const {
        mutateAsync: updateLineItem,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<LineItem>('lineitems', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitLineItem = id ? updateLineItem : createLineItem;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);
        const imageFile = formData.get('image') as File;

        const body: LineItem = {
            name: formData.get('name') as string,
            price: formData.get('price') as unknown as number,
            image: formData.get('image') as string,
        };

        try {
            const response = await submitLineItem(body);

            if (imageFile) {
                const imageData = {
                    imageFile: imageFile,
                    presignedUrl: response.presignedUrl,
                };
                await uploadFileToS3(imageData);
            }

            response.ok ? setShowForm(false) : console.error('failed with response:', response);
        } catch (error) {
            console.error('error creating lineItem', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create LineItem'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>LineItem Form</h1>
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

export default LineItemForm;
