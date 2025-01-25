import {create} from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware';

export type AuthStoreType = {
    token: string;
    username: string;
    email: string;
    is_installed: boolean;

    // Actions
    setToken(token: string): void;
    setUser(user: Pick<AuthStoreType, 'username' | 'email'>): void;
    setInstalledStatus(status: boolean): void;
}

export const useAuthStore = create(persist<AuthStoreType>((set) => ({
    email: '',
    token: '',
    username: '',
    is_installed: false,

    setInstalledStatus: (newStatus) => set({
        is_installed: newStatus,
    }),
    setToken: (token) => set({
        token,
    }),
    setUser: (user) => set({
        username: user.username,
        email: user.email,
    }),
}), {
    name: 'auth_store',
    storage: createJSONStorage(() => localStorage),
}));
