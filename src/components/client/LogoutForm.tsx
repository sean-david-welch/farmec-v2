import utils from '../../styles/Utils.module.css';
import styles from '../../styles/Blogs.module.css';

import { signOutUser } from '../../utils/auth';

const LoginForm = () => {
  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      const idToken = await signOutUser();

      const response = await fetch('/api/auth/logout', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${idToken}` },
      });

      // Handle success (e.g., redirect or show message)
    } catch (error) {
      console.error('Error submitting form:', error);
      // Handle errors
    }
  };

  return (
    <form onSubmit={handleSubmit} class={utils.form}>
      <button type="submit">Logout</button>
    </form>
  );
};

export default LoginForm;
