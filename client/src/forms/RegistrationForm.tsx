import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';
import { useState } from 'react';
import { getRegFields } from '../utils/registrationFields';
import { MachineRegistration } from '../types/miscTypes';
import { useMutateResource } from '../hooks/genericHooks';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
    id?: string;
    registration?: MachineRegistration;
}

const RegistrationForm: React.FC<Props> = ({ id, registration }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = registration ? getRegFields(registration) : getRegFields();

    const {
        mutateAsync: createRegistration,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<MachineRegistration>('registrations');

    const {
        mutateAsync: updateRegistration,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<MachineRegistration>('registrations', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitRegistration = id ? updateRegistration : createRegistration;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const formData = new FormData(event.currentTarget as HTMLFormElement);

        const body: MachineRegistration = {
            dealer_name: formData.get('dealer_name') as string,
            dealer_address: formData.get('dealer_address') as string,
            owner_name: formData.get('owner_name') as string,
            owner_address: formData.get('owner_address') as string,
            machine_model: formData.get('machine_model') as string,
            serial_number: formData.get('serial_number') as string,
            install_date: formData.get('install_date') as string,
            invoice_number: formData.get('invoice_number') as string,
            complete_supply: formData.get('complete_supply') === 'on',
            pdi_complete: formData.get('pdi_complete') === 'on',
            pto_correct: formData.get('pto_correct') === 'on',
            machine_test_run: formData.get('machine_test_run') === 'on',
            safety_induction: formData.get('safety_induction') === 'on',
            operator_handbook: formData.get('operator_handbook') === 'on',
            date: formData.get('date') as string,
            completed_by: formData.get('completed_by') as string,
        };

        try {
            const response = await submitRegistration(body);
            response && !isError && setShowForm(false);
        } catch (error) {
            console.error('Failed to create warranty claim', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create Registration'}
            </button>
            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit}>
                    <h1 className={utils.mainHeading}>Machine Registration Form</h1>
                    {formFields.map(field => (
                        <div key={field.name}>
                            <label htmlFor={field.name}>{field.label}</label>
                            {field.type === 'select' ? (
                                <select name={field.name} id={field.name} defaultValue={field.defaultValue || ''}>
                                    {field.options &&
                                        field.options.map(option => (
                                            <option key={option} value={option}>
                                                {option}
                                            </option>
                                        ))}
                                </select>
                            ) : (
                                <input
                                    type={field.type}
                                    name={field.name}
                                    id={field.name}
                                    placeholder={field.placeholder}
                                    defaultValue={field.defaultValue || ''}
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

export default RegistrationForm;
