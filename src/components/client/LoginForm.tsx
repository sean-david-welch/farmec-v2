import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Blogs.module.css';

import { createSignal } from 'solid-js';

import { signInUser } from '../../utils/auth';

const LoginForm = () => {
  const [email, setEmail] = createSignal('');
  const [password, setPassword] = createSignal('');

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      const idToken = await signInUser(email(), password());

      const response = await fetch('/api/auth/login', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${idToken}` },
      });

      await response.json();
      // Handle success (e.g., redirect or show message)
    } catch (error) {
      console.error('Error submitting form:', error);
      // Handle errors
    }
  };

  return (
    <form onSubmit={handleSubmit} class={utils.form}>
      <label>Email:</label>
      <input type="email" value={email()} onInput={e => setEmail(e.currentTarget.value)} required />

      <label>Password:</label>
      <input type="password" value={password()} onInput={e => setPassword(e.currentTarget.value)} required />

      <button type="submit">Login</button>
    </form>
  );
};

export default LoginForm;
