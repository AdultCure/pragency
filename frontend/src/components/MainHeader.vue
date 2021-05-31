<template>
  <div class="main-header">
    <div class="header-wrapper">
      <router-link :to="{ name: 'Home' }"
        ><div class="header-title-wrapper">
          <img
            src="..\assets\pictures\logo.svg"
            alt="logo"
            class="header-logo"
          />
          <h1 class="header-title">KS Creative</h1>
        </div>
      </router-link>
      <ul class="header-menu">
        <router-link
          :to="{ name: 'Home' }"
          active-class="menu-item-active"
          class="menu-item"
          ><li>Главная</li></router-link
        >
        <router-link
          :to="{ name: 'Services' }"
          active-class="menu-item-active"
          class="menu-item"
          ><li>Услуги</li></router-link
        >
        <router-link
          :to="{ name: 'Orders' }"
          active-class="menu-item-active"
          class="menu-item"
          ><li v-show="$store.state.isAuth">
            Мои заказы
          </li>
        </router-link>
      </ul>
      <div class="header-auth">
        <router-link
          :to="{ name: 'Profile' }"
          active-class="menu-item-active"
          class="header-user-name"
          v-show="$store.state.isAuth"
          >{{ $store.state.currentUser.name }}</router-link
        >
        <button
          class="header-button"
          @click="openModalAuth()"
          v-show="$store.state.isAuth === false"
        >
          Войти
        </button>
        <a href="/"
          ><button
            class="header-button"
            @click="logOut()"
            v-show="$store.state.isAuth"
          >
            Выйти
          </button>
        </a>
      </div>
    </div>
    <modal-auth ref="modalAuth" />
  </div>
</template>

<script>
import ModalAuth from "./modal/ModalAuth.vue";

export default {
  name: "MainHeader",
  components: {
    ModalAuth,
  },
  data() {
    return {};
  },
  methods: {
    focusInput() {
      this.$refs.modalAuth.$refs.loginInput.focus();
    },
    openModalAuth() {
      this.$refs.modalAuth.showShadow = true;
      this.$refs.modalAuth.showAuth = true;
      setTimeout(this.focusInput, 100);
    },
    logOut() {
      localStorage.clear();
      location.reload();
    },
  },
};
</script>

<style lang="scss" scoped>
.login-system {
  position: fixed;
}
.main-header {
  position: sticky;
  top: 0;
  background: #fff;
  padding: 0 40px;
  border-bottom: 1px solid #66ccff;
}
.header-wrapper {
  background: #fff;
  max-width: 1440px;
  width: 100%;
  display: flex;
  margin: 0 auto;
  flex-direction: row;
  height: 79px;
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
.header-menu {
  display: flex;
  align-items: center;
  margin-left: 20px;
  max-width: 300px;
  width: 100%;
  & li {
    cursor: pointer;
    font-size: 14px;
    color: #4d5155;
  }
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
  &:hover {
    border-bottom: 2px solid #59abff;
  }
}
.header-button {
  outline: none;
  transition: linear 0.2s;
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 3px 1px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
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
.menu-item {
  margin-left: 30px;
  height: 20px;
  &:hover {
    border-bottom: 2px solid #59abff;
  }
}
.menu-item-active {
  width: fit-content;
  border-bottom: 2px solid #59abff;
}
</style>
