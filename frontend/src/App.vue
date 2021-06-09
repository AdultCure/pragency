<template>
  <div id="app" class="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
  components: {},
  methods: {},
  mounted() {
    // Проверка авторизации по токену
    if (localStorage.token) {
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
