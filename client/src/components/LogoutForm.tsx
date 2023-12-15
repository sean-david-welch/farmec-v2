import utils from '../../styles/Utils.module.css';

import { signOutUser } from '../../utils/auth';

import { useStore } from '@nanostores/solid';
import { onMount } from 'solid-js';
import { $user, addUser, removeUser } from '../../utils/store';

const LogoutForm = () => {
  const user = useStore($user);

  onMount(() => {
    if (typeof window !== 'undefined') {
      const storedUserData = localStorage.getItem('user');
      if (storedUserData) {
        addUser(JSON.parse(storedUserData));
      }
    }
  });

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    try {
      await signOutUser();

      const response = await fetch('http://localhost:4321/api/auth/logout');

      if (response) {
        const storedUser = localStorage.getItem('user');

        if (storedUser) {
          removeUser();
          localStorage.removeItem('user');
        }

        window.location.href = '/';
      }
    } catch (error) {
      console.error('Error submitting form:', error);
    }
  };

  return user() ? (
    <form onSubmit={handleSubmit} class={utils.form}>
      <button type="submit">Logout</button>
    </form>
  ) : (
    <a href="/login">
      <button type="submit">Login</button>
    </a>
  );
};

export default LogoutForm;
