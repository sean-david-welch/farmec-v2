import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { getFormFields } from '../utils/productFields';
import { uploadFileToS3 } from '../lib/aws';
import { Machine, Product } from '../types/supplierTypes';
import { useMutateResource } from '../hooks/genericHooks';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
    id?: string;
    product?: Product;
    machine: Machine;
}

const ProductForm: React.FC<Props> = ({ id, product, machine }) => {
    const [showForm, setShowForm] = useState(false);

    const formFields = id ? getFormFields(machine, product) : getFormFields(machine);

    const {
        mutateAsync: createProduct,
        isError: isCreateError,
        error: createError,
        isPending: createPending,
    } = useMutateResource<Product>('products');

    const {
        mutateAsync: updateProduct,
        isError: isUpdateError,
        error: updateError,
        isPending: updatingPending,
    } = useMutateResource<Product>('products', id);

    const isError = id ? isUpdateError : isCreateError;
    const error = id ? updateError : createError;

    const submitProduct = id ? updateProduct : createProduct;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);

        const imageFile = formData.get('product_image') as File;

        const body: Product = {
            machine_id: formData.get('machine_id') as string,
            name: formData.get('name') as string,
            product_image: imageFile ? imageFile.name : 'null',
            description: formData.get('description') as string,
            product_link: formData.get('product_link') as string,
        };

        try {
            const response = await submitProduct(body);

            if (imageFile) {
                const productImageData = {
                    imageFile: imageFile,
                    presignedUrl: response.presignedUrl,
                };
                await uploadFileToS3(productImageData);
            }
            response && !isError && setShowForm(false);
        } catch (error) {
            console.error('error creating product', error);
        }
    }

    if (createPending || updatingPending) return <Loading />;
    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create Product'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Product Form</h1>
                    {formFields.map(field => (
                        <div key={field.label}>
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
                                    defaultValue={field.defaultValue}
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

export default ProductForm;
