import { signInWithEmailAndPassword, signOut, getAuth } from '@firebase/auth';

import { initializeApp } from '@firebase/app';

const firebaseConfig = {
  apiKey: import.meta.env.FB_WEB_API_KEY,
  authDomain: import.meta.env.FB_AUTH_URL,
  projectId: import.meta.env.FB_PROJECT_ID,
};

const app = initializeApp(firebaseConfig);

export const auth = getAuth(app);

export const signInUser = async (email: string, password: string): Promise<void> => {
  try {
    const userCredential = await signInWithEmailAndPassword(auth, email, password);
    const user = userCredential.user;
    console.log('User signed in:', user);
  } catch (error: any) {
    console.error('Error signing in:', error);
  }
};

export const signOutUser = async (): Promise<void> => {
  try {
    await signOut(auth);
    console.log('User signed out');
  } catch (error: any) {
    console.error('Error signing out:', error);
  }
};
