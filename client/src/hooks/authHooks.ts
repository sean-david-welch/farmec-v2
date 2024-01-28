import { useEffect } from 'react';

import { auth } from '../lib/auth';
import { updateIsAdmin, useUserStore } from '../lib/store';

const useFirebaseAuthSync = () => {
    const { setIsAuthenticated, setIsAdmin } = useUserStore();

    useEffect(() => {
        const unsubscribe = auth.onAuthStateChanged(async user => {
            if (user) {
                setIsAuthenticated(true);

                const customClaim = (await user.getIdTokenResult()).claims;
                const isAdmin = typeof customClaim.admin === 'boolean' ? customClaim.admin : false;

                console.log(isAdmin);
                updateIsAdmin(isAdmin);
            } else {
                setIsAuthenticated(false);
                setIsAdmin(false);
            }
        });

        return () => unsubscribe();
    }, [setIsAuthenticated, setIsAdmin]);
};

export default useFirebaseAuthSync;
