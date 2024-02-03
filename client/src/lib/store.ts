import { create } from 'zustand';
import { Supplier } from '../types/supplierTypes';

interface UserState {
    isAdmin: boolean;
    setIsAdmin: (isAdmin: boolean) => void;

    isAuthenticated: boolean;
    setIsAuthenticated: (isAuthenticated: boolean) => void;
}

export const useUserStore = create<UserState>()((set) => ({
    isAdmin: false,
    setIsAdmin: (isAdmin: boolean) => set({ isAdmin }),

    isAuthenticated: false,
    setIsAuthenticated: (isAuthenticated: boolean) => set({ isAuthenticated }),
}));

interface SuppliersState {
    suppliers: Supplier[];
    setSuppliers: (suppliers: Supplier[]) => void;
}

export const useSupplierStore = create<SuppliersState>()((set) => ({
    suppliers: [],
    setSuppliers: (suppliers: Supplier[]) => set({ suppliers }),
}));
