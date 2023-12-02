import { initializeApp } from 'firebase/app';
import type { FirebaseApp } from 'firebase/app';

const firebaseConfig = {
  apiKey: 'AIzaSyDd4w_sHjO1X8oo6tafrx0XikIRr82jQtI',
  authDomain: 'farmec-ireland-1675438662747.firebaseapp.com',
  projectId: 'farmec-ireland-1675438662747',
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
