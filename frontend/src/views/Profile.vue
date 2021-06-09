<template>
  <div class="profile">
    <main-header />
    <div class="profile-wrapper">
      <div class="profile-content">
        <div class="profile-card">
          <div class="profile-info">
            <div class="profile-header">
              <img
                src="../assets/pictures/avatar.svg"
                alt="avatar"
                class="profile-avatar"
              />
              <div class="profile-title">
                <h2 class="profile-name">
                  {{ $store.state.currentUser.name }}
                </h2>
                <h3 class="profile-email">
                  {{ $store.state.currentUser.email }}
                </h3>
              </div>
            </div>
          </div>
          <div class="orders-history">
            <h3 class="history-title">История заказов:</h3>
            <div class="orders-card-content" v-if="ordersList !== null">
              <div
                class="history-card"
                v-for="card of ordersList.slice(
                  0,
                  3
                ) /* Выводим только 3 последних заказа */"
                :key="card.id"
              >
                <router-link :to="'/orders/status/' + currentId" class="router"
                  ><div
                    class="history-card-content"
                    @click="currentId = card.id"
                  >
                    <p class="history-card-name">{{ card.category }}</p>
                    <p class="history-card-number">
                      Номер заказа: №{{ card.id }}
                    </p>
                  </div></router-link
                >
                <img
                  src="../assets/pictures/status-icon-1.svg"
                  alt="status"
                  class="history-card-status"
                />
              </div>
            </div>
            <div class="orders-no-cards" v-else>У вас пока нет заказов</div>
            <router-link :to="{ name: 'Orders' }" class="history-look-orders"
              ><p>Посмотреть все заказы</p></router-link
            >
          </div>
        </div>
      </div>
    </div>
    <main-footer />
  </div>
</template>

<script>
import axios from "axios";
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";

export default {
  components: {
    MainHeader,
    MainFooter,
  },
  name: "Profile",
  data() {
    return {
      ordersList: [],
      currentId: "",
    };
  },
  created() {
    // Get-запрос при создании страницы
    axios
      .get("http://localhost:8000/api/order", {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentUser.token}`,
        },
      })
      .then((response) => {
        // Получаем с сервера список всех заказов юзера и кладем в массив
        this.ordersList = response.data.data;
        this.ordersList.reverse();
      })
      .catch((error) => {
        this.loginError = "Упс! Что-то пошло не так :(";
        console.log(error);
      });
  },
};
</script>

<style lang="scss" scoped>
.router {
  color: #4d5155;
}
.profile-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}
.profile-content {
  max-width: 1210px;
  padding: 40px 115px 75px;
}
.profile-card {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  height: 695px;
  padding: 71px 77px 50px 95px;
  display: flex;
  flex-direction: row;
}
.profile-header {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-bottom: 50px;
}
.profile-avatar {
  margin-right: 40px;
}
.profile-name {
  font-weight: 500;
  font-size: 32px;
  line-height: 39px;
  color: #373737;
  margin-bottom: 5px;
  word-break: break-word;
}
.profile-email {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
}
.orders-history {
  margin-left: auto;
  margin-top: 70px;
  max-width: 353px;
  width: 100%;
}
.history-title {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  margin-bottom: 15px;
}
.history-card {
  border: 1px solid #59abff;
  box-sizing: border-box;
  border-radius: 8px;
  max-width: 353px;
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: 8px 15px;
  margin-bottom: 25px;
}
.history-card-name {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  margin: 0;
}
.history-card-number {
  font-weight: normal;
  font-size: 10px;
  line-height: 12px;
  color: #4d5155;
  margin: 5px 0 0 0;
}
.history-card-status {
  width: 20px;
  margin-left: auto;
}
.history-look-orders {
  font-weight: normal;
  font-size: 10px;
  line-height: 12px;
  color: #59abff;
}
.orders-no-cards {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
}
</style>
