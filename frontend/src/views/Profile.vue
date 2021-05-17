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
                <h2 class="profile-name">Имя пользователя</h2>
                <h3 class="profile-email">billyherrington@dungeon.cum</h3>
              </div>
            </div>
            <button
              class="profile-password-button"
              @click="showPassword = !showPassword"
            >
              Изменить пароль
            </button>
            <transition name="fade">
              <form
                class="profile-change-password"
                v-show="showPassword"
                @submit.prevent
              >
                <div class="profile-change">
                  <span class="profile-new-password">Новый пароль: </span
                  ><input type="text" class="profile-password" />
                </div>
                <div class="profile-change">
                  <span class="profile-new-password">Подтвердите пароль: </span
                  ><input type="text" class="profile-password" />
                </div>
                <button class="profile-password-button">Сохранить</button>
              </form>
            </transition>
          </div>
          <div class="orders-history">
            <h3 class="history-title">История заказов:</h3>
            <div class="history-card">
              <div class="history-card-content">
                <p class="history-card-name">{{ $store.state.adList[0] }}</p>
                <p class="history-card-number">Номер заказа: #1234567890</p>
              </div>
              <img
                src="../assets/pictures/status-icon-1.svg"
                alt="status"
                class="history-card-status"
              />
            </div>
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
      showPassword: false,
    };
  },
  methods: {},
};
</script>

<style lang="scss" scoped>
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
}
.profile-email {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
}
.profile-password-button {
  outline: none;
  transition: linear 0.2s;
  background: #fff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 6px;
  max-width: 175px;
  width: 100%;
  height: 37px;
  color: #59abff;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  align-self: flex-end;
  margin: 0 65px 25px 0;
  cursor: pointer;
  &:hover {
    background: #59abff;
    color: #fff;
  }
  &:last-child {
    max-width: 140px;
    width: 100%;
    background: #59abff;
    color: #fff;
  }
}
.profile-new-password {
  display: block;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #373737;
  width: 165px;
}
.profile-password {
  font-size: 14px;
  line-height: 15px;
  color: #4d5155;
  border: none;
  border-bottom: 1px solid #59abff;
  outline: none;
  padding: 0 5px;
  max-width: 200px;
  width: 100%;
  height: 17px;
  margin-bottom: 30px;
}
.profile-change {
  display: flex;
}
.profile-change-password {
  transition: linear 0.2s;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
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
  margin-bottom: 33px;
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
</style>
