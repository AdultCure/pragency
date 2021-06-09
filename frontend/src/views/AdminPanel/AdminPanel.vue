<template>
  <div class="admin-panel">
    <admin-header />
    <div class="admin-wrapper">
      <div class="admin-content">
        <div class="admin-flex">
          <span class="admin-sort">Сортировка заказов:</span>
          <div class="select">
            <select
              class="admin-select"
              size="1"
              name="sort"
              v-model="selectCategory"
            >
              <option value="0">Все категории</option>
              <option
                v-for="category in categories"
                :key="category.id"
                :value="category.category"
                >{{ category.category }}</option
              >
            </select>
          </div>
          <input
            class="admin-search"
            type="seatch"
            name=""
            id=""
            placeholder="Поиск по имени"
            v-model="search"
          />
        </div>
        <div class="card-content" v-if="ordersList !== null">
          <div class="card" v-for="card of filteredOrders" :key="card.id">
            <div class="card-left-content">
              <p class="card-category">{{ card.category }}</p>
              <p class="card-name">{{ card.user_name }}</p>
              <p class="card-date">Статус заказа: {{ card.status }}</p>
              <p class="card-comment">Комментарий:</p>
              <textarea
                disabled
                class="card-comment-text"
                name="comment"
                cols="30"
                rows="10"
                v-model="card.comment"
              ></textarea>
            </div>

            <div class="card-right-content">
              <p class="card-number">Номер заказа: №{{ card.id }}</p>
              <p class="card-date">{{ card.date }}</p>
              <router-link
                :to="'/admin/admin-order/' + currentId"
                class="router"
                ><button class="card-button" @click="currentId = card.id">
                  Подробнее
                </button>
              </router-link>
            </div>
          </div>
        </div>
        <div class="no-card" v-else>Заказы отсутствуют</div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import AdminHeader from "../../components/AdminHeader.vue";
export default {
  name: "AdminPanel",
  components: {
    AdminHeader,
  },

  data() {
    return {
      ordersList: [],
      currentId: "",
      search: "",
      categories: [
        { category: "Реклама в интернете" },
        { category: "Реклама на ТВ" },
        { category: "Рекламные производства" },
        { category: "Наружная Реклама" },
        { category: "Реклама на радио" },
      ],
      selectCategory: "0",
    };
  },
  computed: {
    // Фильтрация заказов
    filteredOrders: function() {
      return (
        this.ordersList
          // Фильтруем по категории
          .filter((card) => {
            return (
              this.selectCategory == "0" || card.category == this.selectCategory
            );
          })

          // Фильтруем по полю поиска
          .filter((card) => {
            return (
              this.search == "" ||
              card.user_name
                .toLowerCase()
                .indexOf(this.search.toLowerCase()) !== -1
            );
          })
      );
    },
  },

  created() {
    // Get-запрос на сервер при создании страницы
    axios
      .get("http://localhost:8000/api/admin", {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentAdmin.token}`, // Заголовок авторизации админа
        },
      })
      .then((response) => {
        // Получаем с сервера список всех заказов и кладем в массив
        this.ordersList = response.data.data;
        this.ordersList.reverse();
      })
      .catch((error) => {
        this.loginError = "Упс! Что-то пошло не так :(";
        console.log(error);
      });
  },
};
</script>

<style lang="scss" scoped>
.admin-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}
.admin-content {
  padding: 48px 115px 62px;
}
.admin-flex {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin-bottom: 8px;
}
.admin-sort {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  margin-right: 15px;
}
.select {
  position: relative;
}
.select select {
  display: block;
  background: none;
  border: none;
  outline: none;
  font-family: inherit;
  font-size: 1rem;
  color: #444;
}
.admin-search {
  display: block;
  margin-left: auto;
  background: #ffffff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  border-radius: 7px;
  width: 300px;
  height: 30px;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  padding: 6px 15px;
  outline: none;
}
.card {
  max-width: 1210px;
  width: 100%;
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  margin-bottom: 40px;
}
p {
  margin: 0 0 5px 0;
}
.card-left-content {
  max-width: 500px;
  width: 100%;
}
.card-comment-text {
  height: 150px;
  background: #f3f3f3;
  border-radius: 8px;
  outline: none;
  border: none;
  resize: none;
  margin: 0;
  padding: 15px;
  width: 100%;
  height: 90px;
}
.card-right-content {
  display: flex;
  flex-direction: column;
}
.router {
  margin-top: auto;
  width: 120px;
}
.card-button {
  align-self: center;
  outline: none;
  transition: linear 0.2s;
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 7px;
  width: 120px;
  height: 35px;
  color: #fff;
  font-size: 14px;
  cursor: pointer;
  &:hover {
    background: #fff;
    color: #59abff;
  }
}
.no-card {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  height: 60px;
  padding-left: 30px;
  display: flex;
  align-items: center;
}
</style>
