import { Routes as ReactRoutes, Route } from 'react-router'
import { lazy } from 'react'

const IndexPage = lazy(() => import('./pages/index'));
const NotFoundPage = lazy(() => import('./pages/404'));

const InstallationIndexPage = lazy(() => import('./pages/installation'));

export const Routes = () => (
    <ReactRoutes>
        {/* Index page */}
        <Route index element={<IndexPage />} />

        {/* Installation */}
        <Route path={'/installation'}>
            {/* Installation index page */}
            <Route index element={<InstallationIndexPage />} />
        </Route>

        {/* 404 Not Found Page */}
        <Route path={'*'} element={<NotFoundPage />} />
    </ReactRoutes>
);