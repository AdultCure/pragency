<template>
  <div class="admin-panel">
    <admin-header />
    <div class="admin-wrapper">
      <div class="admin-content">
        <div class="admin-flex">
          <span class="admin-sort">Сортировка заказов:</span>
          <div class="select">
            <select class="admin-select" size="1" name="sort" id="">
              <option class="admin-select-option" value="Create Data" autofocus
                >Дата создания</option
              >
              <option value="zalupa artema">zalupa artema</option>
            </select>
          </div>
          <input
            class="admin-search"
            type="seatch"
            name=""
            id=""
            placeholder="Поиск"
          />
        </div>
        <div
          class="card"
          v-for="card of $store.state.adminAdList"
          :key="card.id"
        >
          <div class="card-left-content">
            <p class="card-category">{{ card.category }}</p>
            <p class="card-name">{{ card.user_name }}</p>
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
            <router-link :to="{ name: 'AdminOrder' }" class="router"
              ><button
                class="card-button"
                @click="
                  $store.state.selectAd.name = card.user_name;
                  $store.state.selectAd.category = card.category;
                  $store.state.selectAd.date = card.date;
                  $store.state.selectAd.id = card.id;
                  $store.state.selectAd.comment = card.comment;
                "
              >
                Подробнее
              </button>
            </router-link>
          </div>
        </div>
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
  created() {
    axios
      .get("http://localhost:8000/api/admin", {
        headers: {
          Authorization: `Bearer ${this.$store.state.currentUser.token}`,
        },
      })
      .then((response) => {
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
.admin-select {
  height: 17px;
  background: #e3e3e3;
  color: #373737;
  border: none;
  outline: none;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
}
.admin-select-option {
  display: block;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #373737;
  width: 107px;
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
</style>
