import utils from '../styles/Utils.module.css';

import FormDialog from './FormDialog';

import { useState } from 'react';
import { Employee } from '../types/aboutTypes';

import { uploadFileToS3 } from '../lib/aws';
import { useMutateResource } from '../hooks/genericHooks';
import { employeeFormFields } from '../utils/aboutFields';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface Props {
    id?: string;
    employee?: Employee;
}

const EmployeeForm: React.FC<Props> = ({ id, employee }) => {
    const [showForm, setShowForm] = useState(false);
    const formFields = employee ? employeeFormFields(employee) : employeeFormFields();

    const {
        mutateAsync: createEmployee,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<Employee>('employees');

    const {
        mutateAsync: updateEmployee,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<Employee>('employees', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitEmployee = id ? updateEmployee : createEmployee;

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget as HTMLFormElement);
        const imageFile = formData.get('profile_image') as File;

        const body: Employee = {
            name: formData.get('name') as string,
            email: formData.get('email') as string,
            role: formData.get('role') as string,
            profile_image: formData.get('profile_image') as string,
        };

        try {
            const response = await submitEmployee(body);

            if (imageFile) {
                const imageData = {
                    imageFile: imageFile,
                    presignedUrl: response.presignedUrl,
                };
                await uploadFileToS3(imageData);
            }

            response.ok ? setShowForm(false) : console.error('failed with response:', response);
        } catch (error) {
            console.error('error creating employee', error);
        }
    }

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create Employee'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form className={utils.form} onSubmit={handleSubmit} encType="multipart/form-data">
                    <h1 className={utils.mainHeading}>Employee Form</h1>
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

export default EmployeeForm;
