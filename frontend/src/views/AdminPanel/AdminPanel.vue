<template>
  <div class="admin-panel">
    <admin-header />
    <div class="admin-wrapper">
      <div class="admin-content">
        <div class="admin-flex">
          <span class="admin-sort">Сортировка заказов:</span>
          <select class="admin-select" size="1" name="sort" id="">
            <option class="admin-select-option" value="Create Data" autofocus
              >Дата создания</option
            >
            <option value="zalupa artema">zalupa artema</option>
          </select>
          <input
            class="admin-search"
            type="seatch"
            name=""
            id=""
            placeholder="Поиск"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AdminHeader from "../../components/AdminHeader.vue";
export default {
  name: "AdminPanel",
  components: {
    AdminHeader,
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
</style>
