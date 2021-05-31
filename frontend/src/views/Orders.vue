<template>
  <div class="orders">
    <main-header />
    <div class="orders-wrapper">
      <div class="orders-content" v-if="$store.state.userAdList !== null">
        <div
          class="orders-card"
          v-for="userAd of $store.state.userAdList"
          :key="userAd.id"
        >
          <router-link :to="{ name: 'Status' }" class="router">
            <ul
              class="orders-card-content active-card"
              @click="
                $store.state.selectAd.name = userAd.category;
                $store.state.selectAd.data = userAd.date;
                $store.state.selectAd.id = userAd.id;
                $store.state.selectAd.comment = userAd.comment;
              "
            >
              <li class="orders-name">{{ userAd.category }}</li>
              <li class="orders-date">{{ userAd.date }}</li>
              <li class="orders-number">Заказ №{{ userAd.id }}</li>
            </ul>
          </router-link>
        </div>
      </div>
      <div class="orders-no-cards" v-else>
        <h3 class="orders-no-cards-header">
          Тут пока ничего нет. Вы можете заказать услугу
          <router-link :to="{ name: 'Services' }" class="orders-no-cards-link"
            >здесь
          </router-link>
        </h3>
      </div>
    </div>
    <main-footer />
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";
import axios from "axios";

export default {
  components: {
    MainFooter,
    MainHeader,
  },
  name: "Orders",
  data() {
    return {};
  },
  computed: mapGetters(["fullUserAdList"]),
  mounted() {
    axios
      .get("http://localhost:8000/api/order", {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentUser.token}`,
        },
      })
      .then((response) => {
        this.$store.state.userAdList = response.data.data;
        if (this.$store.state.userAdList !== null) {
          this.$store.state.userAdList.reverse();
        }
      })
      .catch((error) => {
        this.loginError = "Упс! Что-то пошло не так :(";
        console.log(error);
      });
  },
};
</script>

<style lang="scss" scoped>
.orders-wrapper {
  min-height: calc(100vh - 220px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}

.router {
  color: #4d5155;
}

.active-card:hover {
  color: #59abff;
}

.orders-content {
  max-width: 1210px;
  padding: 0 115px;
}

.orders-card {
  height: 80px;
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  margin: 60px 0;
}

.orders-card-content {
  display: flex;
  height: 80px;
  align-items: center;
  padding-left: 25px;
  flex-direction: row;
}

.orders-name {
  display: block;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  max-width: 300px;
  width: 100%;
}

.orders-date {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
}

.orders-number {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  margin-left: auto;
  padding-right: 80px;
}

.orders-no-cards {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  height: 60px;
  margin-top: 60px;
  padding-left: 30px;
  display: flex;
  align-items: center;
}

.orders-no-cards-header {
  font-weight: normal;
  font-size: 18px;
  line-height: 29px;
}

.orders-no-cards-link {
  color: #59abff;
}
</style>
