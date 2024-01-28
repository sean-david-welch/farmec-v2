import { useEffect } from 'react';

import { auth } from '../lib/auth';
import { useUserStore } from '../lib/store';

const useFirebaseAuthSync = () => {
    const { setIsAuthenticated, setIsAdmin } = useUserStore();

    console.log('auth hook');

    useEffect(() => {
        const unsubscribe = auth.onAuthStateChanged(user => {
            if (user) {
                // User is signed in, update Zustand store
                setIsAuthenticated(true);
                console.log('user is authenticated', user);
                // Optionally set isAdmin based on your criteria
                // setIsAdmin(checkIfAdmin(user));
            } else {
                // User is signed out, update Zustand store
                console.log('user is signed out', user);
                setIsAuthenticated(false);
                setIsAdmin(false);
            }
        });

        // Cleanup subscription on unmount
        return () => unsubscribe();
    }, [setIsAuthenticated, setIsAdmin]);
};

export default useFirebaseAuthSync;
