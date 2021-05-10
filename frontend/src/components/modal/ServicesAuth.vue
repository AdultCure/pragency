<template>
  <transition name="fade">
    <div
      class="modal-shadow"
      @mousedown.self="closeModal"
      @keydown.esc="closeModal"
      v-show="showShadow"
    >
      <div class="modal" v-show="showModalServices">
        <div class="modal-close" @click="closeModal">&#10006;</div>
        <div class="modal-content">
          <h2 class="modal-header">Вы должны быть авторизированы</h2>
          <p class="modal-text">
            Для оформления заказа вам нужно быть зарегистрированным
            пользователем нашей платформы.
          </p>
          <div class="modal-buttons">
            <button class="modal-login" @click="showModalAuth">Войти</button>
            <button class="modal-cancel" @click="closeModal">Отмена</button>
          </div>
        </div>
      </div>
      <modal-auth ref="modalAuth" @closeServModel="showModalServices = true" />
    </div>
  </transition>
</template>

<script>
import ModalAuth from "./ModalAuth.vue";
export default {
  name: "ServicesAuth",
  components: {
    ModalAuth,
  },
  data() {
    return {
      showShadow: false,
      showModalServices: false,
    };
  },
  methods: {
    closeModal() {
      this.showModalServices = false;
      this.showShadow = false;
    },
    showModalAuth() {
      this.showModalServices = false;
      this.$refs.modalAuth.showShadow = true;
      this.$refs.modalAuth.showAuth = true;
    },
  },
};
</script>

<style lang="scss" scoped>
.modal-shadow {
  position: fixed;
  top: 0;
  left: 0;
  min-height: 100vh;
  width: 100%;
  background: rgba(53, 53, 53, 0.3);
}
.modal {
  max-width: 472px;
  width: 100%;
  display: flex;
  flex-direction: column;
  background: #ffffff;
  box-shadow: 0px 1px 18px 2px rgba(0, 0, 0, 0.15);
  border-radius: 12px;
  align-items: flex-end;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 15px;
}
.modal-close {
  cursor: pointer;
  width: 12px;
  height: 12px;
  color: #a0a0a0;
}
.modal-content {
  align-self: flex-start;
  margin: 30px 40px;
  display: flex;
  flex-direction: column;
}
.modal-header {
  font-weight: 500;
  font-size: 18px;
  text-align: center;
  color: #59abff;
  max-width: 327px;
  margin-bottom: 30px;
}
.modal-text {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  margin-bottom: 32px;
}
.modal-buttons {
  max-width: 260px;
  width: 100%;
  margin: 0;
  align-self: flex-end;
}
.modal-login {
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  border-radius: 4px;
  width: 120px;
  height: 36px;
  color: #fff;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  outline: none;
  transition: linear 0.2s;
  cursor: pointer;
  &:hover {
    background: #fff;
    color: #59abff;
  }
}
.modal-cancel {
  background: #a6a6a6;
  border: 1px solid #a6a6a6;
  box-sizing: border-box;
  border-radius: 4px;
  width: 120px;
  height: 36px;
  color: #fff;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  margin-left: 20px;
  outline: none;
  transition: linear 0.2s;
  cursor: pointer;
  &:hover {
    background: #fff;
    color: #a6a6a6;
  }
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
