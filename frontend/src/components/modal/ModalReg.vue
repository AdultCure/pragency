<template>
  <transition name="fadereg">
    <div class="modal" v-show="showReg">
      <div class="modal-close" @click="$emit('closeRegModal')">
        &#10006;
      </div>
      <div class="modal-content">
        <h3 class="modal-header">Регистрация</h3>
        <form action="" @submit.prevent="submitForm">
          <input
              ref="nameInput"
              id="name"
              v-model="state.name"
              type="text"
              class="modal-input"
              :class="v$.name.$error ? 'form-error-input' : ''"
              placeholder="Ф.И.О"
          />
          <p class="form-error" v-if="v$.name.$error">
            Минимальная длинна ввода 6 символов
          </p>
          <input
              id="email"
              v-model="state.email"
              type.prevent="email"
              class="modal-input"
              :class="v$.email.$error ? 'form-error-input' : ''"
              placeholder="E-mail"
          />
          <p class="form-error" v-if="v$.email.$error">
            Некорректный E-mail
          </p>
          <input
              id="password"
              v-model="state.password.password"
              type="password"
              class="modal-input"
              :class="v$.password.password.$error ? 'form-error-input' : ''"
              placeholder="Пароль"
          />
          <p class="form-error" v-if="v$.password.password.$error">
            Минимальная длинна ввода 6 символов. Только латинские символы
          </p>
          <input
              id="repeatPassword"
              v-model="state.password.confirm"
              type="password"
              class="modal-input"
              :class="v$.password.confirm.$error ? 'form-error-input' : ''"
              placeholder="Повторите пароль"
          />
          <p class="form-error" v-if="v$.password.confirm.$error">
            Пароли не совпадают
          </p>
          <button class="modal-button" type="submit">
            Зарегистрироваться
          </button>
        </form>
      </div>
    </div>
  </transition>
</template>

<script>
import useValidate from "@vuelidate/core";
import {alphaNum, email, minLength, required, sameAs,} from "@vuelidate/validators";
import {computed, reactive} from "vue";
import axios from "axios";

export default {
  name: "ModalReg",
  setup() {
    const state = reactive({
      name: "",
      email: "",
      password: {
        password: "",
        confirm: "",
      },
    });
    const rules = computed(() => {
      return {
        name: {required, minLength: minLength(6)},
        email: {required, email},
        password: {
          password: {required, minLength: minLength(6), alphaNum},
          confirm: {
            required,
            sameAs: sameAs(state.password.password),
            alphaNum,
          },
        },
      };
    });
    const v$ = useValidate(rules, state);

    return {
      state,
      v$,
    };
  },
  data() {
    return {
      showReg: false,
    };
  },

  methods: {
    submitForm() {
      this.v$.$validate();
      if (!this.v$.$error) {
        this.showReg = false;
        this.$emit("closeRegModal");
        this.regPost()
      }
    },

    async regPost() {
      await axios.post(`http://localhost:8000/auth/sign-up`,
          {
            "name": this.state.name.toString(),
            "email": this.state.email.toString(),
            "password": this.state.password.password.toString()
          })
          .then(response => {
            console.log(response);
      })
    }
  },
};
</script>

<style lang="scss" scoped>
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

&
:hover {
  background: #fff;
  color: #59abff;
}

}
.form-error {
  font-size: 12px;
  color: #ff2626;
  margin-top: -25px;
  padding-left: 10px;
}

.form-error-input {
  border-bottom: 1px solid #ff2626;
}

.fadereg-enter-active,
.fadereg-leave-active {
  transition: opacity 0.3s ease;
}

.fadereg-enter-from,
.fadereg-leave-to {
  opacity: 0;
}
</style>
