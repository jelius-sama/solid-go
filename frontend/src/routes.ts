import { type RouteDefinition } from '@solidjs/router';

import ProfileData from '@/pages/profile/profile.data';
import { lazy } from 'solid-js';

export const routes: RouteDefinition[] = [
    {
        path: '/',
        component: lazy(() => import('@/pages/home')),
    },
    {
        path: '/settings',
        component: lazy(() => import('@/pages/settings')),
    },
    {
        path: '/profile',
        component: lazy(() => import('@/pages/profile')),
        preload: ProfileData,
    },
    {
        path: '**',
        component: lazy(() => import('@/pages/not-found')),
    },
];
