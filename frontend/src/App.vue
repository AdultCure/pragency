<template>
  <div id="app" class="app">
    <!-- <notification v-show="$store.state.showNotify" /> -->
    <router-view />
  </div>
</template>

<script>
// import Notification from "./components/notifications/Notification";
export default {
  name: "App",
  components: {
    // Notification,
  },
  methods: {},
  mounted() {
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
