<template>
  <transition name="fade">
    <div
      class="modal-shadow"
      @mousedown.self="closeModal"
      @keydown.esc="closeModal"
      v-show="showShadow"
    >
      <div class="modal" v-show="showAuth">
        <div class="modal-close" @click="closeModal">&#10006;</div>
        <div class="modal-content">
          <h3 class="modal-header">Авторизация</h3>
          <form action="" @submit.prevent="logIn">
            <input
              type="text"
              name="login"
              class="modal-input"
              placeholder="E-mail"
              v-model="email"
              ref="loginInput"
            />
            <input
              type="password"
              class="modal-input"
              placeholder="Пароль"
              v-model="password"
            />
            <p class="modal-description">
              Для авторизации в платформе требуется ввести логин и пароль,
              созданные при регистрации
            </p>
            <p class="modal-ask">Вы еще не зарегистрированы в платформе ?</p>
            <span class="modal-registration" @click="openModalReg"
              >Зарегистрироваться</span
            >
            <button class="modal-button" type="submit">
              Войти
            </button>
          </form>
        </div>
      </div>
      <modal-reg ref="modalReg" @closeRegModal="closeModal" />
    </div>
  </transition>
</template>

<script>
import ModalReg from "./ModalReg.vue";
export default {
  components: {
    ModalReg,
  },
  name: "ModalAuth",
  data() {
    return {
      email: "",
      password: "",
      showAuth: false,
      showShadow: false,
    };
  },
  methods: {
    focusInput() {
      this.$refs.modalReg.$refs.nameInput.focus();
    },
    openModalReg() {
      this.showAuth = false;
      this.$refs.modalReg.showReg = true;
      setTimeout(this.focusInput, 100);
    },
    closeModal() {
      this.showShadow = false;
      this.showAuth = false;
      this.$refs.modalReg.showReg = false;
      this.$emit("closeServModal");
    },
    logIn() {
      this.$store.state.isAuth = true;
      this.showShadow = false;
      this.showAuth = false;
      this.$refs.modalReg.showReg = false;
      this.$emit("closeServModal");
      const authFormData = {
        email: this.email,
        password: this.password,
      };
      console.log(authFormData);
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
  max-width: 320px;
  width: 100%;
  display: flex;
  flex-direction: column;
  background: #ffffff;
  box-shadow: 0px 1px 18px 2px rgba(0, 0, 0, 0.15);
  border-radius: 12px;
  align-items: flex-end;
  position: absolute;
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
  margin: 48px 50px 78px;
  max-width: 220px;
  width: 100%;
}
.modal-header {
  font-weight: 500;
  font-size: 18px;
  line-height: 22px;
  color: #59abff;
  text-align: center;
  margin-bottom: 57px;
}
.modal-input {
  font-size: 14px;
  line-height: 15px;
  color: #4d5155;
  border: none;
  border-bottom: 1px solid #59abff;
  outline: none;
  padding: 0 10px;
  max-width: 200px;
  width: 100%;
  height: 35px;
  margin-bottom: 30px;
}
.modal-description {
  font-size: 12px;
  color: #a0a0a0;
}
.modal-ask {
  color: #a0a0a0;
  font-size: 10px;
}
.modal-registration {
  cursor: pointer;
  display: block;
  font-size: 10px;
  color: #ff2626;
  margin-bottom: 60px;
}
.modal-button {
  transition: linear 0.2s;
  cursor: pointer;
  border: 1px solid #59abff;
  box-sizing: border-box;
  background: #59abff;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 7px;
  max-width: 220px;
  width: 100%;
  height: 43px;
  color: #ffffff;
  font-size: 12px;
  &:hover {
    background: #fff;
    color: #59abff;
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
