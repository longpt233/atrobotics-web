<template>
  <div class="col-2">
    <el-menu
      active-text-color="#dde2ff"
      background-color="#363740"
      class="el-menu-vertical-demo"
      default-active="2"
      text-color="#a4a6b3"
      @open="handleOpen"
      @close="handleClose"
    >
      <div>
        <img class="leading" src="@/assets/images/logo-nobg.jpg" />
        <span class="title">Dashboard</span>
      </div>

      <el-menu-item index="1" v-for="tab in tabList" :key="tab">
        <img class="leading" src="@/assets/images/logo-nobg.jpg" />
        <span class="title">{{ tab }}</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'SideBar',
  data() {
    return {
      tabList: ['Overview', 'Tickets', 'Settings'],
    };
  },
};
</script>

<style lang="scss" scoped>
.leading {
  width: 20px;
  height: auto;
  margin-right: 10px;
}

.title {
  font-size: 19px;
  line-height: 24px;
}
</style>
