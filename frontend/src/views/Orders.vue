<template>
  <div class="orders">
    <main-header />
    <div class="orders-wrapper">
      <div class="orders-content">
        <div class="orders-card-active">
          <div class="orders-card-info">
            <h2 class="orders-name-active">{{ $store.state.adList[0] }}</h2>
            <span class="orders-number-active">Номер заказа №123456789</span>
            <span class="orders-date-active"
              >Дата создание заказа (25.05.2021)</span
            >
            <span class="orders-comment">Ваш комментарий:</span>
            <textarea
              disabled
              class="orders-comment-text"
              name="comment"
              cols="30"
              rows="10"
            ></textarea>
          </div>
          <button class="orders-status">Статус заказа</button>
        </div>
        <div class="orders-card">
          <ul class="orders-card-content">
            <li class="orders-name">{{ $store.state.adList[4] }}</li>
            <li class="orders-date">25.05.2021</li>
            <li class="orders-number">Номер заказа №123456789</li>
          </ul>
        </div>
      </div>
    </div>
    <main-footer />
  </div>
</template>

<script>
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";
export default {
  components: {
    MainFooter,
    MainHeader,
  },
  name: "Orders",
};
</script>

<style lang="scss" scoped>
.orders-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}
.orders-content {
  max-width: 1210px;
  padding: 0 115px;
}
.orders-card-active {
  margin: 50px 0 40px;
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  display: flex;
  flex-direction: row;
  padding: 20px 50px 50px 33px;
  justify-content: space-between;
}
.orders-card-info {
  max-width: 700px;
  width: 100%;
  display: flex;
  flex-direction: column;
}
.orders-name-active {
  font-weight: normal;
  font-size: 24px;
  line-height: 29px;
  color: #59abff;
}
.orders-number-active {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  margin-bottom: 30px;
}
.orders-date-active {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  margin-bottom: 13px;
}
.orders-comment {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #59abff;
  margin-bottom: 6px;
}
.orders-comment-text {
  height: 150px;
  background: #f3f3f3;
  border-radius: 8px;
  outline: none;
  border: none;
  resize: none;
}
.orders-status {
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 3px 1px rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  color: #ffffff;
  max-width: 200px;
  width: 100%;
  height: 31px;
  cursor: pointer;
  transition: linear 0.2s;
  align-self: flex-end;
  margin: 0 0 -23px 23px;
  &:hover {
    background: #fff;
    color: #59abff;
  }
}
.orders-card {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  padding: 30px 25px;
  margin-bottom: 40px;
}
.orders-card-content {
  display: flex;
  flex-direction: row;
}
.orders-name {
  display: block;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  max-width: 300px;
  width: 100%;
}
.orders-date {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
}
.orders-number {
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  margin-left: auto;
  padding-right: 80px;
}
</style>
