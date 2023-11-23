import { atom } from 'nanostores';
import type User from '../types/user';

import { useStore } from '@nanostores/solid';

export const $user = atom<User[]>([]);

export const addUser = (user: User) => {
  $user.set([...$user.get(), user]);
};

export const removeUser = (userUid: string) => {
  $user.set($user.get().filter(user => user.uid !== userUid));
};
