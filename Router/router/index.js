import { createRouter, createWebHistory } from "vue-router";

import HomeCmp from "@/views/HomeScreen"
import AboutCmp from "@/views/AboutScreen";

const routes = [
    {
        name: "HomePage",
        path: "/",
        component: HomeCmp
    },
    {
        name: "AboutPage",
        path: "/about",
        component: AboutCmp
    },
    {
        name: "DetailPage",
        path: "/details/:userID",
        component: () => import("@/views/DetailsScreen")
    }
];

const router = createRouter({
    routes: routes,
    history: createWebHistory()
});

export default router;