import { useState, useEffect } from 'react';
import utils from '../../styles/Utils.module.css';

const RegisterForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('user');

    useEffect(() => {
        if (typeof window !== 'undefined') {
            const storedUserData = localStorage.getItem('user');
            if (storedUserData) {
                // addUser(JSON.parse(storedUserData)); // Update this part based on your global state management strategy
            }
        }
    }, []);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        try {
            const response = await fetch('http://localhost:4321/api/auth/register', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password, role }),
            });

            const result = await response.json();

            if (response.ok) {
                setEmail('');
                setPassword('');
                setRole('user');
                // Handle successful registration (e.g., redirect to login or show a success message)
            } else {
                console.error('Registration failed:', result);
            }
        } catch (error) {
            console.error('Error submitting form:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit} className={utils.form}>
            <label>Email:</label>
            <input type="email" value={email} onChange={e => setEmail(e.target.value)} required />

            <label>Password:</label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)} required />

            <label>Role:</label>
            <select value={role} onChange={e => setRole(e.target.value)} required>
                <option value="user">User</option>
                <option value="admin">Admin</option>
            </select>

            <button type="submit">Register</button>
        </form>
    );
};

export default RegisterForm;
