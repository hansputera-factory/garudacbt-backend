import { Routes as ReactRoutes, Route } from 'react-router'
import { lazy } from 'react'
import { AppLayout } from './layouts/app-layout';

const IndexPage = lazy(() => import('./pages/index'));
const NotFoundPage = lazy(() => import('./pages/404'));

export const Routes = () => (
    <ReactRoutes>
        <Route element={<AppLayout />}>
            {/* Index page */}
            <Route index element={<IndexPage />} />

            {/* 404 Not Found Page */}
            <Route path={'*'} element={<NotFoundPage />} />
        </Route>
    </ReactRoutes>
);