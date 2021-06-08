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
    meta: { auth: true },
    component: () => import("../views/Orders.vue"),
  },
  {
    path: "/orders/status/:id",
    name: "Status",
    meta: { auth: true },
    component: () => import("../views/Status.vue"),
  },
  {
    path: "/profile",
    name: "Profile",
    meta: { auth: true },
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
    meta: { admin: true },
    component: () => import("../views/AdminPanel/AdminPanel.vue"),
  },
  {
    path: "/admin/admin-order/:id",
    name: "AdminOrder",
    meta: { admin: true },
    component: () => import("../views/AdminPanel/AdminOrder.vue"),
  },
  {
    path: "/error",
    name: "Error",
    component: () => import("../views/Error.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from, next) => {
  const currentUser = localStorage.token;
  const requireAuth = to.matched.some((record) => record.meta.auth);
  if (requireAuth && !currentUser) {
    next({ name: "Error" });
  } else {
    next();
  }
});
router.beforeEach((to, from, next) => {
  const currentAdmin = localStorage.admin;
  const requireAdmin = to.matched.some((record) => record.meta.admin);
  if (requireAdmin && !currentAdmin) {
    next({ name: "Error" });
  } else {
    next();
  }
});

export default router;
