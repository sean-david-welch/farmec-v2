import utils from '../../styles/Utils.module.css';

import { $user, addUser } from '../../utils/userStore';
import { useStore } from '@nanostores/solid';
import { createSignal, onMount } from 'solid-js';

import type User from '../../types/user';

const RegisterForm = () => {
  const user = useStore($user);

  console.log('User', user);

  const [email, setEmail] = createSignal('');
  const [password, setPassword] = createSignal('');
  const [role, setRole] = createSignal('user');

  const fetchUserData = () => {
    if (typeof window !== 'undefined') {
      const storedUserData = localStorage.getItem('user');
      if (storedUserData) {
        addUser(JSON.parse(storedUserData));
      }
    }
  };

  onMount(() => {
    fetchUserData();
    console.log('Mounted User:', user());
  });

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      const response = await fetch('http://localhost:4321/api/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          email: email(),
          password: password(),
          role: role(),
        }),
      });

      const result = await response.json();

      if (response.ok) {
        setEmail('');
        setPassword('');
        setRole('user');
      } else {
        console.error('Registration failed:', result);
      }
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  };

  return user() && user()?.role === 'admin' ? (
    <form onSubmit={handleSubmit} class={utils.form}>
      <label>Email:</label>
      <input type="email" value={email()} onInput={e => setEmail(e.currentTarget.value)} required />

      <label>Password:</label>
      <input type="password" value={password()} onInput={e => setPassword(e.currentTarget.value)} required />

      <label>Role:</label>
      <select value={role()} onChange={e => setRole(e.currentTarget.value)} required>
        <option value="user">User</option>
        <option value="admin">Admin</option>
      </select>

      <button type="submit">Register</button>
    </form>
  ) : null;
};

export default RegisterForm;
