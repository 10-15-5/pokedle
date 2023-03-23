import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import ClassicView from '../views/ClassicView.vue';
import FlavorTextView from '../views/FlavortextView.vue';
import SilhouetteView from '../views/SilhuetView.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
        },
        {
            path: '/classic',
            name: 'classic',
            component: ClassicView,
            // route level code-splitting
            // this generates a separate chunk (About.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            //component: () => import("../views/AboutView.vue"),
        },
        {
            path: '/flavortext',
            name: 'flavortext',
            component: FlavorTextView,
        },
        {
            path: '/silhouette',
            name: 'silhouette',
            component: SilhouetteView,
        },
    ],
});

export default router;
