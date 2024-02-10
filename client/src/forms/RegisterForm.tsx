import utils from '../styles/Utils.module.css';
import FormDialog from './FormDialog';

import { useState } from 'react';
import { UserData } from '../types/dataTypes';
import { useMutateResource } from '../hooks/genericHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons/faPenToSquare';

interface Props {
    id?: string;
}

const RegisterForm: React.FC<Props> = ({ id }) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('user');
    const [showForm, setShowForm] = useState(false);

    const {
        mutateAsync: createUser,
        isError: isCreateError,
        error: createError,
    } = useMutateResource<UserData>('users');

    const {
        mutateAsync: updateUser,
        isError: isUpdateError,
        error: updateError,
    } = useMutateResource<UserData>('users', id);

    const error = id ? updateError : createError;
    const isError = id ? isUpdateError : isCreateError;
    const submitUser = id ? updateUser : createUser;

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const body: UserData = {
            email: email,
            password: password,
            role: role,
        };

        try {
            const response = await submitUser(body);
            response.ok ? setShowForm(false) : console.error('failed with response:', response);
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    };

    return (
        <section id="form">
            <button className={utils.btnForm} onClick={() => setShowForm(!showForm)}>
                {id ? <FontAwesomeIcon icon={faPenToSquare} /> : 'Create User'}
            </button>

            <FormDialog visible={showForm} onClose={() => setShowForm(false)}>
                <form onSubmit={handleSubmit} className={utils.form}>
                    <label>Email:</label>
                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />

                    <label>Password:</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />

                    <label>Role:</label>
                    <select value={role} onChange={(e) => setRole(e.target.value)} required>
                        <option value="user">User</option>
                        <option value="admin">Admin</option>
                    </select>

                    <button className={utils.btnForm} type="submit">
                        Register
                    </button>
                </form>

                {isError && <p>Error: {error?.message}</p>}
            </FormDialog>
        </section>
    );
};

export default RegisterForm;
