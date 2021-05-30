<template>
  <div class="orders">
    <main-header/>
    <div class="orders-wrapper">
      <div class="orders-content" v-if="responseData.length">
        <div
            class="orders-card"
            v-for="userAd of fullUserAdList"
            :key="userAd.id"
        >
          <router-link :to="{ name: 'Status' }" class="router"
          >
            <ul
                class="orders-card-content active-card"
                @click="
                $store.state.selectAd.name = userAd.name;
                $store.state.selectAd.data = userAd.data;
                $store.state.selectAd.id = userAd.id;
                $store.state.selectAd.comment = userAd.comment;
              "
            >
              <li class="orders-name">{{ userAd.name }}</li>
              <li class="orders-date">{{ userAd.data }}</li>
              <li class="orders-number">{{ userAd.id }}</li>
            </ul>
          </router-link
          >
        </div>
      </div>
      <div class="orders-no-cards" v-else>
        <h3 class="orders-no-cards-header">
          Тут пока ничего нет. Вы можете заказать услугу
          <router-link :to="{ name: 'Services' }" class="orders-no-cards-link"
          >здесь
          </router-link
          >
        </h3>
      </div>
    </div>
    <main-footer/>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
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
    return {
      responseData: []
    };
  },
  computed: mapGetters(["fullUserAdList"]),
  mounted() {
    axios.get("http://localhost:8000/api/order",
        {headers: {Authorization: `Bearer ${this.$store.state.currentUser.token}`}}
    )
        .then((response) => {
          response.data = this.responseData
          console.log(response.data)
        })
        .catch((error) => {
          this.loginError = "Упс! Что-то пошло не так :(";
          console.log(error);
        })
  }
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
