<template>
  <transition name="fade">
    <div class="notification">
      <div class="card">
        <div class="message">{{ $store.state.notymessage }}</div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "Notification",
  data() {
    return {};
  },
  mounted() {},
};
</script>

<style lang="scss" scoped>
.notification {
  position: fixed;
  top: 85%;
  left: 82%;
}
.card {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 280px;
  height: 80px;
  background: #59abff;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 7px;
  color: #fff;
}
.message {
  font-size: 14px;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
