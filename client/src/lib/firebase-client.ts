import { initializeApp } from 'firebase/app';
import type { FirebaseApp } from 'firebase/app';
import config from '../utils/env';

const firebaseConfig = {
    apiKey: config.firebaseApiKey,
    authDomain: config.firebaseAuthDomain,
    projectId: config.firebaseProjectId,
};

if (!firebaseConfig.apiKey || !firebaseConfig.authDomain || !firebaseConfig.projectId) {
    console.error('Firebase configuration is missing');
}

let app: FirebaseApp;
try {
    app = initializeApp(firebaseConfig);
} catch (error) {
    console.error('Error initializing Firebase:', error);
}

export { app };
