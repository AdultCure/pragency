<template>
  <transition name="fade">
    <div
      class="modal-shadow"
      @mousedown.self="closeModal"
      @keydown.esc="closeModal"
      v-show="showShadow"
    >
      <div class="modal" v-show="showModalOrder">
        <div class="modal-close" @click="closeModal">&#10006;</div>
        <div class="modal-content">
          <p class="modal-text">
            Чтобы заказать услугу оставьте комментарий, вскором времени с вами
            свяжется менеджер.<br />
            Заказ можно отслеживать в “Мои заказы”
          </p>
          <h3 class="modal-comment">Комментарий:</h3>
          <form action="" class="modal-form">
            <textarea
              class="modal-input"
              name="comment"
              maxlength="2000"
              v-model="comment"
              required
            />
            <span v-if="showAlert" class="comment-alert"
              >Пожалуйста, оставьте комментарий к заказу</span
            >
            <div class="modal-buttons">
              <button class="modal-order" @click.prevent="create">
                Заказать
              </button>
              <button class="modal-cancel" @click.prevent="closeModal">
                Отмена
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "ServicesOrder",
  data() {
    return {
      showShadow: false,
      showModalOrder: false,
      showAlert: false,
      comment: "",
    };
  },
  methods: {
    closeModal() {
      this.showModalOrder = false;
      this.showShadow = false;
    },
    create() {
      // Проверяем на наличие комментария: если есть - передаем функцию в Services и там вызываем submit(), иначе показываем ошибку
      if (this.comment.length) {
        this.$emit("submitHandler");
        this.closeModal();
      } else {
        this.showAlert = true;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.comment-alert {
  font-size: 15px;
  color: #ffa740;
}
.modal-shadow {
  position: fixed;
  top: 0;
  left: 0;
  min-height: 100vh;
  width: 100%;
  background: rgba(53, 53, 53, 0.3);
}
.modal {
  max-width: 1000px;
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
  margin: 0 0 25px 85px;
}
.modal-text {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  color: #4d5155;
  margin: 23px 0 25px;
  max-width: 800px;
}
.modal-comment {
  margin-bottom: 5px;
  font-weight: normal;
  font-size: 18px;
  color: #353535;
}
.modal-form {
  display: flex;
  flex-direction: column;
  max-width: 800px;
}
.modal-input {
  background: #f3f3f3;
  border-radius: 8px;
  outline: none;
  border: none;
  height: 204px;
  padding: 10px;
  resize: none;
  margin-bottom: 25px;
}
.modal-buttons {
  max-width: 260px;
  width: 100%;
  margin: 0;
  align-self: flex-end;
}
.modal-order {
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
