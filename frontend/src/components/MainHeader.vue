<template>
  <div class="main-header">
    <div class="header-wrapper">
      <img
        src="..\assets\pictures\logo.svg"
        alt="logo"
        class="header-logo"
        @click="$router.push('/')"
      />
      <h1 class="header-title">KS Creative</h1>
      <ul class="header-menu">
        <li class="menu-main" @click="$router.push('/')">Главная</li>
        <li class="menu-services" @click="$router.push('/services')">Услуги</li>
      </ul>
      <button class="header-button">Войти</button>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
.main-header {
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
  height: 80px;
  align-items: center;
}
.header-logo {
  width: 35px;
  height: 30px;
  cursor: pointer;
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
}
</style>
