import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Blogs.module.css';

import { createSignal } from 'solid-js';

import { signInUser } from '../../utils/auth';

const LoginForm = () => {
  const [email, setEmail] = createSignal('');
  const [password, setPassword] = createSignal('');
  const [errorMessage, setErrorMessage] = createSignal('');

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      const idToken = await signInUser(email(), password());

      const response = await fetch('/api/auth/login', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${idToken}` },
      });

      const user = await response.json();

      if (user) {
        setEmail('');
        setPassword('');

        localStorage.setItem('user', JSON.stringify(user));

        window.location.reload();
      }
    } catch (error: any) {
      if (error.response && error.response.status === 401) {
        setErrorMessage('Incorrect email or password.');
      } else {
        setErrorMessage('An unexpected error occurred. Please try again later.');
      }

      if (error instanceof Error) {
        console.error('Error submitting form:', error.message);
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} class={utils.form}>
      <label>Email:</label>
      <input
        type="email"
        value={email()}
        onInput={e => {
          setEmail(e.currentTarget.value);
          setErrorMessage('');
        }}
        required
      />

      <label>Password:</label>
      <input
        type="password"
        value={password()}
        onInput={e => {
          setPassword(e.currentTarget.value);
          setErrorMessage('');
        }}
        required
      />

      {errorMessage() && <div class={styles.errorMessage}>{errorMessage()}</div>}
      <button type="submit">Login</button>
    </form>
  );
};

export default LoginForm;
