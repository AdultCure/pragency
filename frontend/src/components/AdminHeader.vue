<template>
  <div class="admin-header">
    <div class="header-wrapper">
      <div class="header-title-wrapper">
        <router-link :to="{ name: 'AdminPanel' }"
          ><img
            src="..\assets\pictures\logo.svg"
            alt="logo"
            class="header-logo"
          />

          <h1 class="header-title">KS Creative</h1>
        </router-link>
      </div>
      <div class="header-auth">
        <span class="header-user-name">ADMIN</span>
        <button class="header-button" @click="logOut()">
          Выйти
        </button>
      </div>
    </div>
    <notification ref="notify" v-show="$store.state.showNotify" />
  </div>
</template>

<script>
import Notification from "./notifications/Notification.vue";
export default {
  name: "AdminHeader",
  components: {
    Notification,
  },
  data() {
    return {};
  },
  methods: {
    logOut() {
      localStorage.clear();
      this.$store.state.isAdmin = false;
      this.$store.state.isAuth = false;
      this.$router.push("/");
    },
  },
};
</script>

<style lang="scss" scoped>
a {
  display: flex;
  flex-direction: row;
}
.admin-header {
  position: sticky;
  top: 0;
  background: #fff;
  padding: 0 40px;
}
.header-wrapper {
  background: #fff;
  max-width: 1440px;
  width: 100%;
  display: flex;
  margin: 0 auto;
  flex-direction: row;
  height: 70px;
  align-items: center;
}
.header-title-wrapper {
  display: flex;
  cursor: pointer;
}
.header-logo {
  width: 35px;
  height: 30px;
}
.header-title {
  font-weight: normal;
  font-size: 24px;
  color: #59abff;
  padding-bottom: 10px;
  margin-left: 5px;
}
.header-auth {
  margin-left: auto;
}
.header-user-name {
  margin-right: 30px;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  cursor: pointer;
}
.header-button {
  outline: none;
  transition: linear 0.2s;
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 7px;
  width: 79px;
  height: 31px;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  &:hover {
    background: #fff;
    color: #59abff;
  }
}
</style>
