<template>
  <div class="container">
    <div class="d-flex justify-content-between">
      <div class="d-flex align-items-center">
        <img class="button me-5" src="@/assets/images/logo-nobg.jpg" />
        <div class="me-4 button button-home">Trang chủ</div>
        <div class="me-4 button button-products">Sản phẩm</div>
        <div class="me-4 button button-blog">Fanpage</div>
      </div>

      <div class="d-flex align-items-center">
        <el-input v-model="input3" placeholder="Please input" class="input-with-select">
          <template #append>
            <el-button type="primary">
              <el-icon> <IconSearch /> </el-icon>
            </el-button>
          </template>
        </el-input>
      </div>
    </div>
  </div>
</template>

<script>
import { Search as IconSearch } from '@element-plus/icons-vue';
export default {
  components: { IconSearch },
};
</script>

<style lang="scss" scoped>
* {
  font-size: 16px;
  line-height: 20px;
  font-weight: bold;
}

img {
  height: 50px;
  width: auto;
}

.button {
  cursor: pointer;
}
</style>
