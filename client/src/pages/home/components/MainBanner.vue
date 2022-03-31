<template>
  <div>
    <el-carousel trigger="click" height="764px" interval="0" indicator-position="none">
      <el-carousel-item v-for="item in 4" :key="item">
        <img src="@/assets/images/banner-home.png" />
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {}
</script>

<style lang="scss" scoped>

</style>
