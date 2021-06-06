<template>
  <div class="admin-order">
    <admin-header />
    <div class="order-wrapper">
      <div class="order-content">
        <div class="order-card">
          <div class="card-left-content">
            <p class="order-info">
              Категория услуги: {{ $store.state.selectAd.category }}
            </p>
            <p class="order-info">
              Ф.И.О. клиента: {{ $store.state.selectAd.name }}
            </p>
            <p class="order-comment">Комментарий:</p>
            <textarea
              disabled
              class="order-comment-text"
              name="comment"
              cols="30"
              rows="10"
              v-model="$store.state.selectAd.comment"
            ></textarea>
          </div>
          <div class="card-right-content">
            <p class="order-info">
              Номер заказа: №{{ $store.state.selectAd.id }}
            </p>
            <p class="order-info">
              Дата создания: {{ $store.state.selectAd.date }}
            </p>
            <p class="order-info-change">Изменить статус заказа:</p>
            <div class="select">
              <select class="admin-select" size="1" name="sort" id="">
                <option
                  class="admin-select-option"
                  value="Create Data"
                  autofocus
                  >Дата создания</option
                >
                <option value="zalupa artema">zalupa artema</option>
              </select>
            </div>
            <div class="order-buttons">
              <button class="order-save">Сохранить</button>
              <button class="order-delete">Удалить</button>
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
      id: this.$store.state.selectAd.id,
      data: "",
    };
  },
  created() {
    axios
      .get(`http://localhost:8000/api/order/${this.id}`, {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentUser.token}`,
        },
      })
      .then((response) => {
        this.data = response.data.category;
        this.$store.state.adminAdList = response.data.data;
        if (this.$store.state.adminAdList !== null) {
          this.$store.state.adminAdList.reverse();
        }
      })
      .catch((error) => {
        this.loginError = "Упс! Что-то пошло не так :(";
        console.log(error);
      });
  },
};
</script>

<style lang="scss" scoped>
.order-wrapper {
  min-height: calc(100vh - 220px);
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
