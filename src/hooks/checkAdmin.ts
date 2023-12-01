import { createSignal } from 'solid-js';

export const useFetchUserData = () => {
  const [isAdmin, setAdmin] = createSignal(false);

  const fetchUserData = () => {
    if (typeof window !== 'undefined') {
      const userString = localStorage.getItem('user');
      if (!userString) {
        console.log('User not found in localStorage');
        return;
      }

      const user = JSON.parse(userString);
      if (!user || !user.uid) {
        return;
      }

      const role = user.role;

      setAdmin(role || false);
    }
  };

  return { isAdmin, fetchUserData };
};
