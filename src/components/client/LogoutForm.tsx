import utils from '../../styles/Utils.module.css';

import { signOutUser } from '../../utils/auth';
import { removeUser } from '../../utils/userStore';

const LogoutForm = () => {
  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      const idToken = await signOutUser();

      const response = await fetch('http://localhost:4321/api/auth/logout', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json', Authorization: `Bearer ${idToken}` },
      });

      if (response) {
        const storedUser = localStorage.getItem('user');

        if (storedUser) {
          const userData = JSON.parse(storedUser);
          removeUser(userData.uid);
          localStorage.removeItem('user');
        }

        window.location.reload();
      }
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit} class={utils.form}>
      <button type="submit">Logout</button>
    </form>
  );
};

export default LogoutForm;
