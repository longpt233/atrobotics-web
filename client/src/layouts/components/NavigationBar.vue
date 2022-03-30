<template>
  <div class="container">
    <div class="d-flex justify-centent-between">
      <div class="d-flex align-items-center">
        <img src="@/assets/images/logo-nobg.jpg" class="me-5" />
        <div class="me-4 button button-home">Home</div>
        <div class="me-4 button button-pages">Pages</div>
        <div class="me-4 button button-products">Products</div>
        <div class="me-4 button button-blog">Blog</div>
        <div class="me-4 button button-shop">Shop</div>
      </div>

      <div class="d-flex align-items-center">
        <div class="mt-4">
          <el-input v-model="input3" placeholder="Please input" class="input-with-select">
            <template #prepend>
              <el-select v-model="select" placeholder="Select" style="width: 110px">
                <el-option label="Restaurant" value="1" />
                <el-option label="Order No." value="2" />
                <el-option label="Tel" value="3" />
              </el-select>
            </template>
            <template #append>
              <el-button :icon="Search" />
            </template>
          </el-input>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  components: {},
}
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
