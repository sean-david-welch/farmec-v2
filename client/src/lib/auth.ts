import { app } from './firebase';
import { getAuth, signOut, signInWithEmailAndPassword, browserSessionPersistence } from 'firebase/auth';
import { useUserStore } from './store';

export const auth = getAuth(app);

export const signInUser = async (email: string, password: string): Promise<string | undefined> => {
    let idToken: string | undefined;

    auth.setPersistence(browserSessionPersistence);

    try {
        const userCredential = await signInWithEmailAndPassword(auth, email, password);

        const user = userCredential.user;

        idToken = await user.getIdToken();

        const customClaim = (await user.getIdTokenResult()).claims;
        const isAdmin = typeof customClaim.admin === 'boolean' ? customClaim.admin : false;

        const setUserStore = useUserStore.getState();

        setUserStore.setIsAdmin(isAdmin);
        setUserStore.setIsAuthenticated(true);
    } catch (error: any) {
        console.error('Error signing in:', error);
    }

    return idToken;
};

export const signOutUser = async (): Promise<void> => {
    try {
        await signOut(auth);

        const setUserStore = useUserStore.getState();

        setUserStore.setIsAdmin(false);
        setUserStore.setIsAuthenticated(false);
    } catch (error: any) {
        console.error('Error signing out:', error);
    }
};
