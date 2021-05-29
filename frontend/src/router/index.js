import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/services",
    name: "Services",
    component: () => import("../views/Services.vue"),
  },
  {
    path: "/orders",
    name: "Orders",
    component: () => import("../views/Orders.vue"),
  },
  {
    path: "/orders/status",
    name: "Status",
    component: () => import("../views/Status.vue"),
  },
  {
    path: "/profile",
    name: "Profile",
    component: () => import("../views/Profile.vue"),
  },
  {
    path: "/admin",
    name: "AdminLogin",
    component: () => import("../views/AdminPanel/AdminLogin.vue"),
  },
  {
    path: "/admin/admin-panel",
    name: "AdminPanel",
    component: () => import("../views/AdminPanel/AdminPanel.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
