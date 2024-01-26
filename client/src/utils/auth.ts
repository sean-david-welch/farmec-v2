import { app } from '../lib/firebase-client';
import { getAuth, signOut, signInWithEmailAndPassword, browserSessionPersistence } from 'firebase/auth';

const auth = getAuth(app);

export const signInUser = async (email: string, password: string): Promise<string | undefined> => {
    let idToken: string | undefined;

    auth.setPersistence(browserSessionPersistence);

    try {
        const userCredential = await signInWithEmailAndPassword(auth, email, password);

        const user = userCredential.user;

        idToken = await user.getIdToken();
    } catch (error: any) {
        console.error('Error signing in:', error);
    }

    return idToken;
};

export const signOutUser = async (): Promise<void> => {
    try {
        await signOut(auth);
    } catch (error: any) {
        console.error('Error signing out:', error);
    }
};
