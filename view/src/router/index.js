import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    path: "/",
    component: () => import("@/page/layout.vue"),
    redirect: "/home",
    children: [
      {
        path: "home",
        name: "Home",
        component: () => import("@/page/home/index.vue"),
      },
    ],
  },
  { path: "/about", name: "about", component: () => import("@/page/about/index.vue") },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
