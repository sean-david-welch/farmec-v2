import { atom } from 'nanostores';
import type User from '../types/user';

export const $user = atom<User | null>(null);

export const addUser = (userData: User) => {
    $user.set(userData);
};

export const removeUser = () => {
    $user.set(null);
};
