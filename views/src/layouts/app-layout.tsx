import useSWR from 'swr';
import { useAuthStore } from "@/stores/auth-store"
import React from 'react';
import { APIRoutes } from '@/utils/api-routes';
import { ResponseType } from '@/types/response';

import { Outlet } from 'react-router';

export const AppLayout = () => {
    const authStore = useAuthStore();
    const {isLoading, data, mutate} = useSWR(APIRoutes.routes.install.check, ([url, method]) => APIRoutes.do<ResponseType<{}>>(url, method));

    React.useEffect(() => {
        console.log(data);
        if (data?.ok) {
            authStore.setInstalledStatus(true);
        } else {
            authStore.setInstalledStatus(false);
        }
    }, [isLoading]);

    React.useEffect(() => {
        APIRoutes.token = authStore.token;
        mutate();
    }, [authStore.token]);

    return (
        <section>
                
                {!isLoading && !authStore.is_installed && (
                    <p>
                        todo..
                    </p>
                )}

                {!isLoading && authStore.is_installed && (
                    <Outlet />
                )}
        </section>
    )
}