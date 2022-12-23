import { createRouter, createWebHistory} from "vue-router";

import HomeCmp from "@/views/HomeScreen";
import AuthorCmp from "@/views/AuthorScreen";
import HelloWorld from "@/components/HelloWorld";

const routes = [
    {
        name: "HomePage",
        path:"/socket",
        component:HomeCmp
    },
    {
        name: "AuthorsPage",
        path:"/authors",
        component: AuthorCmp

    },
    {
        name: "BarPage",
        path: "/bar",
        component: HelloWorld
    }
];

const router = createRouter({
    routes: routes,
    history: createWebHistory()
});

export default router;