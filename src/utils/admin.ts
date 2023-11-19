import { getAuth } from 'firebase-admin/auth';

export const verifyToken = async (token: string) => {
  if (!token) {
    throw new Error('No token provided');
  }

  try {
    const decodedToken = await getAuth().verifyIdToken(token);
    return decodedToken;
  } catch (error) {
    throw new Error('Invalid token');
  }
};
