<template>
  <div class="services">
    <main-header />
    <div class="services-wrapper">
      <h2 class="services-header">Услуги</h2>
      <div class="services-content">
        <div class="services-card">
          <div class="services-card-title">
            <img
              src="../assets/pictures/adinternet.svg"
              alt="internet"
              class="services-card-img"
            />
            <h2 class="services-card-header">Реклама в интернете</h2>
          </div>
          <p class="services-card-text">
            Предлагает рекламодателям достаточно широкий спектр рекламных услуг.
            Продвижение в социальных сетях, вирусный маркетинг.
          </p>
          <button class="services-card-button">Заказать</button>
        </div>
        <div class="services-card">
          <div class="services-card-title">
            <img
              src="../assets/pictures/adtv.svg"
              alt="internet"
              class="services-card-img"
            />
            <h2 class="services-card-header">Реклама на ТВ</h2>
          </div>
          <p class="services-card-text">
            Реклама на ТВ будет лучше работать и приносить прибыль только в том
            случае, если реклама вашей организации или фирмы будет очень часто
            появляться на экранах телевизоров.
          </p>
          <button class="services-card-button">Заказать</button>
        </div>
        <div class="services-card">
          <div class="services-card-title">
            <img
              src="../assets/pictures/adproduction.svg"
              alt="internet"
              class="services-card-img"
            />
            <h2 class="services-card-header">Рекламные производства</h2>
          </div>
          <p class="services-card-text">
            Печать рекламы на авто, действенный способ рассказать о своих
            услугах или товарах широкому кругу потенциальных потребителей.
          </p>
          <button class="services-card-button">Заказать</button>
        </div>
        <div class="services-card">
          <div class="services-card-title">
            <img
              src="../assets/pictures/adout.svg"
              alt="internet"
              class="services-card-img"
            />
            <h2 class="services-card-header">Наружная Реклама</h2>
          </div>
          <p class="services-card-text">
            Наружная реклама направлена на привлечение потенциальных покупателей
            при помощи различных видов носителей. Как и в любом другом важном
            деле
          </p>
          <button class="services-card-button">Заказать</button>
        </div>
        <div class="services-card">
          <div class="services-card-title">
            <img
              src="../assets/pictures/adradio.svg"
              alt="internet"
              class="services-card-img"
            />
            <h2 class="services-card-header">Реклама на радио</h2>
          </div>
          <p class="services-card-text">
            Большая часть рекламы на радио состоит из аудиороликов, длительность
            которых может варьироваться от пятнадцати секунд до минуты.
          </p>
          <button class="services-card-button">Заказать</button>
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
  name: "Services",
  components: {
    MainHeader,
    MainFooter,
  },
};
</script>

<style lang="scss" scoped>
.services-wrapper {
  min-height: calc(100vh - 80px - 80px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}
.services-content {
  padding: 0 100px 70px 100px;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(600px, 1fr));
  grid-gap: 40px;
}
.services-header {
  font-weight: normal;
  font-size: 48px;
  line-height: 59px;
  margin: 0 auto;
  width: 170px;
  padding: 30px 0;
}
.services-card {
  display: flex;
  flex-direction: column;
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
}
.services-card-title {
  display: flex;
  flex-direction: row;
  align-items: center;
  margin: 30px 40px;
}
.services-card-header {
  font-weight: normal;
  font-size: 36px;
  line-height: 44px;
  margin-left: 30px;
}
.services-card-text {
  height: 55px;
  margin: 0 40px 23px;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
}
.services-card-button {
  background: #59abff;
  border: 1px solid #59abff;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 6px;
  max-width: 175px;
  width: 100%;
  height: 37px;
  color: #ffffff;
  font-weight: normal;
  font-size: 14px;
  line-height: 17px;
  align-self: flex-end;
  margin: 0 65px 25px 0;
  cursor: pointer;
}
</style>
