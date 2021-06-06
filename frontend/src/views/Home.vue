<template>
  <div class="home">
    <main-header />
    <div class="home-wrapper">
      <div class="home-content">
        <section class="about">
          <h2 class="about-header">
            Коротко о нас
          </h2>
          <div class="about-content">
            <ul class="about-text">
              <p>
                Рекламное агентство KS Creative предоставляет клиентам полный
                комплекс рекламных услуг:
              </p>
              <br />
              <li>
                1. &nbsp; Стратегию маркетингового развития. Создают стратегию
                построения и продажи марки для целевой аудитории.
              </li>
              <br />
              <li>
                2. &nbsp; Создание рекламы. Наши арт-директора, которые и
                создают саму рекламу и копирайтеры, которые составляют рекламные
                тексты, способствующие увеличению объемов продаж.
              </li>
              <br />
              <li>
                3. &nbsp; Медиа планирование. Проводит исследования с тем, чтобы
                определить целевую аудиторию для данной рекламной кампании.
              </li>
              <br />
              <li>
                4. &nbsp; Планирование. Определение количественных и
                качественных факторов, затрагивающих восприятие и отношение
                потребителя к рекламируемому товару.
              </li>
              <br />
              <li>
                5. &nbsp; Производство. Это сам процесс физического
                преобразования рекламной стратегии в реальную рекламу в печати,
                рекламу по радио и телевидению, или любой другой вид рекламы.
              </li>
            </ul>

            <img
              src="../assets/pictures/main-picture.svg"
              alt="picture"
              class="about-picture"
            />
          </div>
        </section>
        <section class="working">
          <h2 class="working-header">Как мы работаем</h2>
          <div class="working-content">
            <div class="working-step">
              <img
                src="../assets/pictures/working-1.svg"
                alt="1"
                class="working-img"
              />
              <p class="working-text">
                Зарегистрируйтесь в нашей платформе, прежде чем заказать услугу
              </p>
            </div>
            <div class="working-step">
              <img
                src="../assets/pictures/working-2.svg"
                alt="2"
                class="working-img"
              />
              <p class="working-text">
                Менеджер свяжется с вами и уточнит детали заказа
              </p>
            </div>
            <div class="working-step">
              <img
                src="../assets/pictures/working-3.svg"
                alt="3"
                class="working-img"
              />
              <p class="working-text">
                После обсуждения подготавливаем коммерческое предложение
              </p>
            </div>
            <div class="working-step">
              <img
                src="../assets/pictures/working-4.svg"
                alt="4"
                class="working-img"
              />
              <p class="working-text">
                Принимаем решение о работе, согласовываем и подписываем договор
              </p>
            </div>
            <div class="working-step">
              <img
                src="../assets/pictures/working-5.svg"
                alt="5"
                class="working-img"
              />
              <p class="working-text">
                Предоставляем отчет о проделанных работах
              </p>
            </div>
          </div>
        </section>
        <section class="benefits">
          <h2 class="benefits-header">Что это дает?</h2>
          <div class="benefits-content">
            <div class="benefits-card">
              <img
                src="../assets/pictures/benefits-card-1.svg"
                alt="card-1"
                class="benefits-img"
              />
              <p class="benefits-text">Повышение лояльности покупателей</p>
            </div>
            <div class="benefits-card">
              <img
                src="../assets/pictures/benefits-card-2.svg"
                alt="card-2"
                class="benefits-img"
              />
              <p class="benefits-text">Стимулирование продаж</p>
            </div>
            <div class="benefits-card">
              <img
                src="../assets/pictures/benefits-card-3.svg"
                alt="card-3"
                class="benefits-img"
              />
              <p class="benefits-text">
                Привлечение внимания к продукту или категории товаров
              </p>
            </div>
            <div class="benefits-card">
              <img
                src="../assets/pictures/benefits-card-4.svg"
                alt="card-4"
                class="benefits-img"
              />
              <p class="benefits-text">
                Формирование портрета аудитории бренда
              </p>
            </div>
          </div>
        </section>
      </div>
    </div>
    <main-footer />
  </div>
</template>

<script>
import MainFooter from "../components/MainFooter.vue";
import MainHeader from "../components/MainHeader.vue";
export default {
  name: "Home",
  components: {
    MainHeader,
    MainFooter,
  },
  data() {
    return {
      show: true,
    };
  },
};
</script>

<style lang="scss" scoped>
// HOME
.home {
  background-color: #e3e3e3;
  margin: 0 auto;
}
.home-wrapper {
  min-height: calc(100vh - 220px);
  max-width: 1440px;
  width: 100%;
  margin: 0 auto;
}
.home-content {
  padding: 0 115px;
}
// ABOUT
.about {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  margin-top: 50px;
}
.about-header {
  font-weight: 500;
  font-size: 24px;
  color: #4d5155;
  margin: 40px auto 30px;
  width: 177px;
}
.about-content {
  display: flex;
  align-items: flex-start;
  margin: 0 77px 60px 60px;
  & img {
    padding-top: 14px;
    max-width: 250px;
    width: 100%;
  }
}
.about-text {
  margin-right: 84px;
  font-weight: 500;
  font-size: 14px;
  line-height: 17px;
  & p {
    margin: 0;
  }
}
// WORKING
.working {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  margin-top: 50px;
}
.working-header {
  width: 219px;
  margin: 40px auto 50px;
  font-weight: 500;
  font-size: 24px;
}
.working-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(213px, 1fr));
  grid-gap: 10px;
  align-self: center;
  margin: 0 60px 80px 40px;
}
.working-step {
  text-align: center;
  font-weight: 500;
  font-size: 14px;
  line-height: 17px;
}
// BENEFITS
.benefits {
  margin-bottom: 95px;
}
.benefits-header {
  font-weight: 500;
  font-size: 24px;
  margin: 50px auto 25px;
  width: 168px;
}
.benefits-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  grid-gap: 70px;
}
.benefits-card {
  background: #ffffff;
  border: 1px solid #d0d0d0;
  box-sizing: border-box;
  box-shadow: 0px 1px 8px 1px rgba(0, 0, 0, 0.1);
  border-radius: 13px;
  text-align: center;
  height: 250px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  & img {
    padding-bottom: 20px;
  }
  & p {
    font-weight: 500;
    font-size: 14px;
    line-height: 17px;
    padding: 0 20px;
  }
}
</style>
