import { createStore } from "vuex";
import user from "./user";

export default createStore({
  state: {
    isAuth: false,
    adList: [
      "Реклама в интернете",
      "Реклама на ТВ",
      "Рекламные производства",
      "Наружная Реклама",
      "Реклама на радио",
    ],
  },
  mutations: {},
  actions: {},
  modules: { user },
});
