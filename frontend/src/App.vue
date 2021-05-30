<template>
  <div id="app" class="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
  mounted() {
    if (localStorage.password) {
      this.$store.state.isAuth = true;
    } else {
      this.$store.state.isAuth = false;
    }
  },
};
</script>

<style>
@import url("./app.css");
</style>
