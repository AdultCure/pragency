import { createStore } from "vuex";

export default createStore({
  state: {
    isAuth: false,
    isAdmin: false,
    showNotify: false,
    notymessage: "",
    currentUser: {
      name: localStorage.name,
      token: localStorage.token,
      email: localStorage.email,
    },
    adList: [
      {
        name: "Реклама в интернете",
        description:
          "Предлагает рекламодателям достаточно широкий спектр рекламных услуг. Продвижение в социальных сетях, вирусный маркетинг.",
        img: "adinternet.svg",
      },
      {
        name: "Реклама на ТВ",
        description:
          "Реклама на ТВ будет лучше работать и приносить прибыль только в том случае, если реклама вашей организации или фирмы будет очень часто появляться на экранах телевизоров.",
        img: "adtv.svg",
      },
      {
        name: "Рекламные производства",
        description:
          "Печать рекламы на авто, действенный способ рассказать о своих услугах или товарах широкому кругу потенциальных потребителей.",
        img: "adproduction.svg",
      },
      {
        name: "Наружная Реклама",
        description:
          "Наружная реклама направлена на привлечение потенциальных покупателей при помощи различных видов носителей. Как и в любом другом важном деле",
        img: "adout.svg",
      },
      {
        name: "Реклама на радио",
        description:
          "Большая часть рекламы на радио состоит из аудиороликов, длительность которых может варьироваться от пятнадцати секунд до минуты.",
        img: "adradio.svg",
      },
    ],
    userAdList: [],
    selectAd: {
      name: "",
      data: "",
      id: "",
      comment: "",
    },
  },
  getters: {
    fullAdList(state) {
      return state.adList;
    },
  },
  actions: {},
});
