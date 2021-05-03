<template>
  <div class="main-header">
    <div class="header-wrapper">
      <div class="header-title-wrapper" @click="$router.push('/')">
        <img src="..\assets\pictures\logo.svg" alt="logo" class="header-logo" />
        <h1 class="header-title">KS Creative</h1>
      </div>
      <ul class="header-menu">
        <li class="menu-item" @click="$router.push('/')">Главная</li>
        <li class="menu-item" @click="$router.push('/services')">Услуги</li>
      </ul>
      <button class="header-button" @click="showModal">
        Войти
      </button>
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
  methods: {
    showModal() {
      this.$refs.modalAuth.show = true;
    },
  },
};
</script>

<style lang="scss" scoped>
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
  margin-left: 50px;
  & li {
    cursor: pointer;
    font-size: 14px;
    color: #4d5155;
  }
  & li:last-child {
    margin-left: 30px;
  }
}
.header-button {
  outline: none;
  transition: linear 0.2s;
  margin-left: auto;
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
  height: 20px;
  &:hover {
    border-bottom: 2px solid #59abff;
  }
}
</style>
