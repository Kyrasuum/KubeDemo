import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/about",
        component: () => import("/src/pages/About.vue"),
    },
    {
        path: "/:catchAll(.*)*",
        component: () => import("/src/pages/Landing.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
