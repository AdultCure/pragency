<template>
  <div class="services">
    <main-header />
    <div class="services-wrapper">
      <h2 class="services-header">Услуги</h2>
      <div class="services-content">
        <div class="services-card" v-for="ad of fullAdList" :key="ad.name">
          <div class="services-card-title">
            <img
              v-bind:src="getImgUrl(ad.img)"
              alt="picture"
              class="services-card-img"
            />
            <h2 class="services-card-header">{{ ad.name }}</h2>
          </div>
          <p class="services-card-text">
            {{ ad.description }}
          </p>
          <button
            class="services-card-button"
            @click="
              selectAd.name = ad.name;
              selectAd.description = ad.description;
              openModal();
            "
          >
            Заказать
          </button>
        </div>
      </div>
    </div>
    <main-footer />
    <services-auth ref="servauth" />
    <services-order ref="servorder" @submitHandler="submit()" />
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";
import ServicesAuth from "../components/modal/ServicesAuth.vue";
import ServicesOrder from "../components/modal/ServicesOrder.vue";
import axios from "axios";

export default {
  name: "Services",
  components: {
    MainHeader,
    MainFooter,
    ServicesAuth,
    ServicesOrder,
  },
  data() {
    return {
      selectAd: {
        name: "",
        description: "",
        comment: "",
      },
    };
  },
  computed: mapGetters(["fullAdList"]), // Получаем весь список рекламы из vuex store
  methods: {
    async submit() {
      // Функция заказа из ServicesOrder
      let Data = new Date(); // Создание текущей даты
      let Year = Data.getFullYear();
      let Month = Data.getMonth();
      let Day = Data.getDate();

      await axios // Post-запрос на backend
        .post(
          "http://localhost:8000/api/order",
          // Передаём данные выбранного заказа, дату создания и комментарий из ServicesOrder на сервер
          {
            category: this.selectAd.name.toString(),
            status: "Ваш заказ создан",
            date: `${Day}.${Month + 1}.${Year}`,
            comment: this.$refs.servorder.comment.toString(),
            user_id: this.$store.state.currentUser.id,
            user_name: this.$store.state.currentUser.name,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.currentUser.token}`,
            },
          }
        )
        .then((response) => {
          // Если нет ошибки, показываем сообщение о заказе и пушим на страницу orders
          setTimeout(() => {
            this.$store.state.notymessage = "Заказ принят в обработку";
            this.$store.state.showNotify = true;
            setTimeout(() => (this.$store.state.showNotify = false), 2000);
          });
          this.$router.push("orders");
          console.log(response);
        })
        .catch((error) => {
          // В случае ошибки - выводим в консоль
          this.loginError = "Упс! Что-то пошло не так :(";
          console.log(error);
        });
    },
    openModal() {
      // Проверка авторизации: если авторизирован - открыть попап с заказом, если нет - открыть окно авторизации
      if (this.$store.state.isAuth === false) {
        this.$refs.servauth.showShadow = true;
        this.$refs.servauth.showModalServices = true;
      } else {
        this.$refs.servorder.showShadow = true;
        this.$refs.servorder.showModalOrder = true;
      }
    },
    getImgUrl(pic) {
      // Получение правильного url для изображений
      return require("../assets/pictures/" + pic);
    },
  },
};
</script>

<style lang="scss" scoped>
.services-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}

.services-content {
  padding: 0 100px 70px 100px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(600px, 1fr));
  grid-gap: 40px;
}

.services-header {
  font-weight: normal;
  font-size: 48px;
  line-height: 59px;
  margin: 0 auto;
  width: 170px;
  padding: 30px 0;
}

.services-card {
  display: flex;
  flex-direction: column;
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
}

.services-card-title {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin: 30px 40px;
}

.services-card-header {
  font-weight: normal;
  font-size: 36px;
  line-height: 44px;
  margin-left: 30px;
}

.services-card-image {
  width: 150px;
  height: 150px;
}

.services-card-text {
  height: 55px;
  margin: 0 40px 23px;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
}

.services-card-button {
  outline: none;
  transition: linear 0.2s;
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 6px;
  max-width: 175px;
  width: 100%;
  height: 37px;
  color: #ffffff;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  align-self: flex-end;
  margin: 0 65px 25px 0;
  cursor: pointer;

  &:hover {
    background: #fff;
    color: #59abff;
  }
}
</style>
