import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import Loading from '../layouts/Loading';

import { useState } from 'react';
import { getWarrantyFields, getPartFields } from '../utils/warrantyFields';
import { WarrantyParts, WarrantyClaim } from '../types/miscTypes';
import { useMutateResource } from '../hooks/genericHooks';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
    id?: string;
    warranty?: WarrantyClaim;
}

const WarrantyForm: React.FC<Props> = ({ id, warranty }) => {
    const [showForm, setShowForm] = useState(false);
    const [parts, setParts] = useState([
        {
            part_number: '',
            quantity_needed: '',
            invoice_number: '',
            description: '',
        },
    ]);
    const addPart = () => {
        setParts([
            ...parts,
            {
                part_number: '',
                quantity_needed: '',
                invoice_number: '',
                description: '',
            },
        ]);
    };

    const formFields = warranty ? getWarrantyFields(warranty) : getWarrantyFields();

    const {
        mutateAsync: createWarranty,
        isError: isCreateError,
        error: createError,
        isPending: createPending,
    } = useMutateResource<WarrantyParts>('warranty');

    const {
        mutateAsync: updateWarranty,
        isError: isUpdateError,
        error: updateError,
        isPending: updatingPending,
    } = useMutateResource<WarrantyParts>('warranty', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitWarranty = id ? updateWarranty : createWarranty;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const formData = new FormData(event.currentTarget as HTMLFormElement);

        const partsRequired = parts.map((_, index) => {
            return {
                part_number: event.currentTarget[`part_number_${index}`].value,
                quantity_needed: event.currentTarget[`quantity_needed_${index}`].value,
                invoice_number: event.currentTarget[`invoice_number_${index}`].value,
                description: event.currentTarget[`part_description_${index}`].value,
            };
        });

        const body: WarrantyParts = {
            warranty: {
                dealer: formData.get('dealer') as string,
                dealer_contact: formData.get('dealer_contact') as string,
                owner_name: formData.get('owner_name') as string,
                machine_model: formData.get('machine_model') as string,
                serial_number: formData.get('serial_number') as string,
                install_date: formData.get('install_date') as string,
                failure_date: formData.get('failure_date') as string,
                repair_date: formData.get('repair_date') as string,
                failure_details: formData.get('failure_details') as string,
                repair_details: formData.get('repair_details') as string,
                labour_hours: formData.get('labour_hours') as string,
                completed_by: formData.get('completed_by') as string,
            },
            parts: partsRequired,
        };

        try {
            const response = await submitWarranty(body);
            response && !isError && setShowForm(false);
        } catch (error) {
            console.error('Failed to create wwarranty claim', error);
        }
        setShowForm(false);
    }

    if (createPending || updatingPending) return <Loading />;
    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? (
                    <FontAwesomeIcon icon={faPenToSquare} />
                ) : (
                    <div>
                        Warranty Claim
                        <FontAwesomeIcon icon={faPenToSquare} />
                    </div>
                )}
            </button>
            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit}>
                    <h1 className={utils.mainHeading}>Warranty Claim Form</h1>
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
                    {parts.map((part, index) => {
                        const partFields = getPartFields(part, index);
                        return (
                            <div key={index}>
                                {partFields.map(field => (
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
                            </div>
                        );
                    })}

                    <button type="button" className={utils.btnForm} onClick={addPart}>
                        Add Part
                    </button>

                    <button className={utils.btnForm} type="submit">
                        Submit
                    </button>
                </form>

                {isError && <p>Error: {error?.message}</p>}
            </FormDialog>
        </section>
    );
};

export default WarrantyForm;
