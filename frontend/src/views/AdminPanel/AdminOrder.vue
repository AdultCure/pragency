<template>
  <div class="admin-order">
    <admin-header />
    <div class="order-wrapper">
      <div class="order-content">
        <div class="order-card">
          <div class="card-left-content">
            <p class="order-info">Категория услуги: {{ response.category }}</p>
            <p class="order-info">Ф.И.О. клиента: {{ response.user_name }}</p>
            <p class="order-comment">Комментарий:</p>
            <textarea
              disabled
              class="order-comment-text"
              name="comment"
              cols="30"
              rows="10"
              v-model="response.comment"
            ></textarea>
          </div>
          <div class="card-right-content">
            <p class="order-info">Номер заказа: №{{ response.id }}</p>
            <p class="order-info">Дата создания: {{ response.date }}</p>
            <p class="order-info-change">Изменить статус заказа:</p>
            <div class="select">
              <select
                class="admin-select"
                name="sort"
                id=""
                v-model="response.status"
              >
                <option
                  class="admin-select-option"
                  v-for="option of statusMap"
                  :key="option.id"
                  v-bind:selected="response.status"
                >
                  {{ option.name }}
                </option>
              </select>
            </div>
            <div class="order-buttons">
              <button class="order-save" @click="editOrder">Сохранить</button>
              <button class="order-delete" @click="deleteOrder">Удалить</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import AdminHeader from "@/components/AdminHeader.vue";

export default {
  name: "Status",
  components: {
    AdminHeader,
  },
  data() {
    return {
      id: this.$route.params.id,
      response: {
        id: "",
        category: "",
        status: "",
        date: "",
        comment: "",
        user_id: "",
        user_name: "",
      },
      statusMap: [
        {
          name: "Ваш заказ создан",
          id: "1",
        },
        {
          name: "Заказ принят",
          id: "2",
        },
        {
          name: "Заказ ожидает оплаты",
          id: "3",
        },
        {
          name: "Принимаем решение о работе",
          id: "4",
        },
        {
          name: "Сделка завершена",
          id: "5",
        },
      ],
    };
  },
  created() {
    // Get-запрос на сервер при создании страницы
    axios
      .get(`http://localhost:8000/api/admin/` + `${this.id}`, {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentAdmin.token}`,
        },
      })
      .then((response) => {
        // Получаем данные заказа с сервера и кладем их в объект response
        this.response.id = response.data.id;
        this.response.category = response.data.category;
        this.response.status = response.data.status;
        this.response.date = response.data.date;
        this.response.comment = response.data.comment;
        this.response.user_id = response.data.user_id;
        this.response.user_name = response.data.user_name;
        console.log(this.response.status);
      })
      .catch((error) => {
        console.log(error);
      });
  },
  methods: {
    deleteOrder() {
      // Delete-запрос на сервер
      axios
        .delete(`http://localhost:8000/api/admin/` + `${this.id}`, {
          headers: {
            Authorization: `Bearer ${this.$store.state.currentAdmin.token}`,
          },
        })
        .then((response) => {
          // Удаляем заказ, выводим сообщение и пушим в админ-панель
          setTimeout(() => {
            this.$store.state.notymessage = "Заказ удалён";
            this.$store.state.showNotify = true;
            setTimeout(() => (this.$store.state.showNotify = false), 2000);
          });
          this.$router.push(`/admin/admin-panel`);
          console.log(response.data);
        })
        .catch((error) => {
          console.log(error);
        });
    },

    editOrder() {
      // Put-запрос на сервер
      axios
        .put(
          `http://localhost:8000/api/admin/` + `${this.id}`,
          // Передаем данные о заказе из объекта response на сервер
          {
            category: this.response.category,
            status: this.response.status,
            date: this.response.date,
            comment: this.response.comment,
          },
          {
            headers: {
              Authorization: `Bearer ${this.$store.state.currentAdmin.token}`,
            },
          }
        )
        .then((response) => {
          // При успешном запросе выводим сообщение и пушим в админ-панель
          setTimeout(() => {
            this.$store.state.notymessage = "Изменения сохранены";
            this.$store.state.showNotify = true;
            setTimeout(() => (this.$store.state.showNotify = false), 2000);
          });
          this.$router.push(`/admin/admin-panel`);
          console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style lang="scss" scoped>
.order-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}

.order-content {
  max-width: 1210px;
  padding: 50px 115px 90px;
}

.order-card {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  padding: 40px;
  display: flex;
  justify-content: space-between;
  height: 634px;
}

.card-left-content {
  max-width: 700px;
  width: 100%;
}

.order-info {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  color: #4d5155;
  margin: 0 0 25px 0;
  align-self: flex-end;
}

.order-comment {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  color: #4d5155;
  margin: 0 0 15px 0;
}

.order-comment-text {
  height: 300px;
  max-width: 670px;
  width: 100%;
  background: #f3f3f3;
  border-radius: 8px;
  outline: none;
  border: none;
  resize: none;
  margin: 0;
  padding: 15px;
}

.card-right-content {
  display: flex;
  flex-direction: column;
  max-width: 370px;
  width: 100%;
}

.order-info-change {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  color: #4d5155;
  margin: 80px 0 10px 0;
}

.select {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  border: none;
  outline: none;
}

.admin-select {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  border: none;
  outline: none;
}

.order-buttons {
  margin-top: auto;
}

.order-save {
  width: 160px;
  height: 34px;
  font-weight: normal;
  font-size: 13px;
  line-height: 16px;
  color: #ffffff;
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 3px 1px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  outline: none;
  transition: linear 0.2s;
  cursor: pointer;
  margin-right: 23px;

  &:hover {
    background: #fff;
    color: #59abff;
  }
}

.order-delete {
  width: 160px;
  height: 34px;
  font-weight: normal;
  font-size: 13px;
  line-height: 16px;
  color: #ffffff;
  background: #ff3131;
  border: 1px solid #ff3131;
  box-sizing: border-box;
  box-shadow: 0px 1px 3px 1px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  outline: none;
  transition: linear 0.2s;
  cursor: pointer;
  margin-right: 23px;

  &:hover {
    background: #fff;
    color: #ff3131;
  }
}
</style>
