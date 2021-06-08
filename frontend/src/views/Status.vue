<template>
  <div class="status">
    <main-header />
    <div class="status-wrapper">
      <div class="status-content">
        <div class="status-card">
          <h2 class="status-header">Статус заказа</h2>

          <div class="status-items">
            <div class="status-item">
              <img
                src="../assets/pictures/status-items/status-1.svg"
                alt="status-icon"
                class="status-img"
              />
              <div
                class="status-check"
                :class="{ active: response.status == 'Ваш заказ создан' }"
              ></div>
              <p class="status-text">Ваш заказ создан</p>
            </div>
            <div class="status-item">
              <img
                src="../assets/pictures/status-items/status-2.svg"
                alt="status-icon"
                class="status-img"
              />
              <div
                class="status-check"
                :class="{ active: response.status == 'Заказ принят' }"
              ></div>
              <p class="status-text">Ваш заказ принят</p>
            </div>
            <div class="status-item">
              <img
                src="../assets/pictures/status-items/status-3.svg"
                alt="status-icon"
                class="status-img"
              />
              <div
                class="status-check"
                :class="{ active: response.status == 'Заказ ожидает оплаты' }"
              ></div>
              <p class="status-text">Заказ ожидает оплаты</p>
            </div>
            <div class="status-item">
              <img
                src="../assets/pictures/status-items/status-4.svg"
                alt="status-icon"
                class="status-img"
              />
              <div
                class="status-check"
                :class="{
                  active: response.status == 'Принимаем решение о работе',
                }"
              ></div>
              <p class="status-text">Принимаем решение о работе</p>
            </div>
            <div class="status-item">
              <img
                src="../assets/pictures/status-items/status-5.svg"
                alt="status-icon"
                class="status-img"
              />
              <div
                class="status-check"
                :class="{ active: response.status == 'Сделка завершена' }"
              ></div>
              <p class="status-text">Сделка завершена</p>
            </div>
          </div>
          <div class="status-line"></div>
          <div class="status-info">
            <h2 class="status-name">{{ response.category }}</h2>
            <span class="status-number">Заказ №{{ response.id }}</span>
            <span class="status-date">Дата создания: {{ response.data }}</span>
            <span class="status-comment">Ваш комментарий:</span>
            <textarea
              disabled
              class="status-comment-text"
              name="comment"
              cols="30"
              rows="10"
              v-model="response.comment"
            ></textarea>
          </div>
        </div>
      </div>
    </div>
    <main-footer />
  </div>
</template>

<script>
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";
import axios from "axios";
export default {
  name: "Status",
  components: {
    MainFooter,
    MainHeader,
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
    };
  },
  created() {
    axios
      .get(`http://localhost:8000/api/order/` + `${this.id}`, {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentUser.token}`,
        },
      })
      .then((response) => {
        this.response.id = response.data.id;
        this.response.category = response.data.category;
        this.response.status = response.data.status;
        this.response.date = response.data.date;
        this.response.comment = response.data.comment;
        this.response.user_id = response.data.user_id;
        this.response.user_name = response.data.user_name;
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>

<style lang="scss" scoped>
.status-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}
.status-content {
  max-width: 1210px;
  padding: 50px 115px 90px;
}
.status-card {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  padding: 45px 53px 68px;
}
.status-header {
  margin-bottom: 50px;
  text-align: center;
  font-weight: normal;
  font-size: 28px;
  line-height: 34px;
  text-align: center;
  color: #59abff;
}
.status-line {
  border: 1px solid #59abff;
  margin-top: -115px;
  margin-bottom: 115px;
}
.status-items {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 60px;
}
.status-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.status-img {
  width: 50px;
  height: 50px;
}
.status-check {
  width: 15px;
  height: 15px;
  border: 1px solid #727272;
  border-radius: 50%;
  box-sizing: border-box;
  margin: 26px 0 12px;
}
.active {
  background-color: #9cf58d;
  border-color: #9cf58d;
}
.status-text {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #4d5155;
  margin: 0;
}
.status-info {
  max-width: 700px;
  width: 100%;
  display: flex;
  flex-direction: column;
}
.status-name {
  font-weight: normal;
  font-size: 24px;
  line-height: 29px;
  color: #59abff;
  margin-bottom: 5px;
}
.status-number {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  margin-bottom: 30px;
}
.status-date {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  color: #4d5155;
  margin-bottom: 13px;
}
.status-comment {
  font-weight: normal;
  font-size: 18px;
  line-height: 22px;
  color: #59abff;
  margin-bottom: 6px;
}
.status-comment-text {
  height: 150px;
  background: #f3f3f3;
  border-radius: 8px;
  outline: none;
  border: none;
  resize: none;
  margin: 0;
  padding: 15px;
}
</style>
