import useSWR from 'swr';
import { useAuthStore } from "@/stores/auth-store"
import React from 'react';
import { Skeleton } from '@/components/ui/skeleton';
import { Card } from '@/components/ui/card';
import { APIRoutes } from '@/lib/api-routes';
import { ResponseType } from '@/types/response';

import InstallationIndexPage from '@/pages/installation';
import { Outlet } from 'react-router';

export const AppLayout = () => {
    const authStore = useAuthStore();
    const {isLoading, data, mutate} = useSWR<ResponseType<{}>>([APIRoutes.routes.install.check, undefined, false], APIRoutes.do);

    React.useEffect(() => {
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
                {isLoading && (
                    <div className="flex h-screen items-center justify-center">
                        <Card className="w-[300px] space-y-4 p-6">
                            <Skeleton className="h-12 w-12 rounded-full mx-auto" />
                            <Skeleton className="h-6 w-3/4 mx-auto" />
                            <Skeleton className="h-4 w-1/2 mx-auto" />
                            <Skeleton className="h-4 w-full" />
                            <Skeleton className="h-4 w-5/6" />
                        </Card>
                    </div>
                )}

                {!isLoading && !authStore.is_installed && (
                    <InstallationIndexPage />
                )}

                {!isLoading && authStore.is_installed && (
                    <Outlet />
                )}
        </section>
    )
}