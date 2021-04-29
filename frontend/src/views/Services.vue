<template>
  <div class="services">
    <main-header />
    <div class="services-wrapper">
      <h2>Услуги</h2>
    </div>
    <main-footer />
  </div>
</template>

<script>
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";
export default {
  name: "Services",
  components: {
    MainHeader,
    MainFooter,
  },
};
</script>

<style lang="scss" scoped>
.services-wrapper {
  min-height: calc(100vh - 80px - 80px);
}
</style>
