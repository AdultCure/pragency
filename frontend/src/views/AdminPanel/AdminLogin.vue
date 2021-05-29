<template>
  <div class="admin">
    <div class="modal">
      <div class="modal-content">
        <h3 class="modal-header">Вход в панель администратора</h3>
        <form action="" @submit.prevent="formSubmit">
          <p class="input-header">Логин:</p>
          <input type="text" name="login" class="modal-input" v-model="login" />
          <p class="input-header">Пароль:</p>
          <input type="password" class="modal-input" v-model="password" />
          <button class="modal-button" type="submit">
            Войти
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AdminLogin",
  data() {
    return {
      login: "",
      password: "",
    };
  },
  methods: {
    formSubmit() {
      this.signIn();
    },
    signIn() {
      this.$store.state.isAdmin = true;
      this.$router.push("/admin/admin-panel");
    },
  },
};
</script>

<style lang="scss" scoped>
.admin {
  background: #ffffff;
  min-height: 100vh;
}
.modal {
  width: 320px;
  height: 421px;
  background: #6f6f6f;
  box-shadow: 0px 5px 10px 2px rgba(0, 0, 0, 0.15);
  border-radius: 12px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
.modal-content {
  margin: 60px 50px;
}
.modal-header {
  font-size: 18px;
  line-height: 22px;
  text-align: center;
  color: #ffffff;
  font-weight: 500;
  margin-bottom: 24px;
}
.input-header {
  margin: 0 0 5px 0;
  font-weight: 500;
  font-size: 12px;
  line-height: 15px;
  color: #ffffff;
}
.modal-input {
  width: 220px;
  height: 30px;
  margin-bottom: 20px;
  outline: none;
  border: none;
}
.modal-button {
  background: #ffffff;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 7px;
  font-weight: 500;
  font-size: 12px;
  line-height: 15px;
  color: #000;
  width: 220px;
  height: 43px;
  outline: none;
  border: none;
  cursor: pointer;
  &:hover {
    background: #e5e5e5;
  }
}
</style>
