import styles from '../../styles/Blogs.module.css';
import utils from '../../styles/Utils.module.css';
import { createSignal } from 'solid-js';

const RegisterForm = () => {
  const [email, setEmail] = createSignal('');
  const [password, setPassword] = createSignal('');
  const [role, setRole] = createSignal('user');

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      const response = await fetch('/api/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          email: email(),
          password: password(),
          role: role(),
        }),
      });

      const result = await response.json();
      console.log(result);
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

      <label>Role:</label>
      <select value={role()} onChange={e => setRole(e.currentTarget.value)} required>
        <option value="user">User</option>
        <option value="admin">Admin</option>
      </select>

      <button type="submit">Register</button>
    </form>
  );
};

export default RegisterForm;
