<template>
  <div class="container-fluid">
    <div
      class="heading-bar-content container d-flex justify-content-between align-items-center"
    >
      <div>
        <span class="email-address">
          <img src="@/assets/images/icon-email.svg" />
          <a class="router-link" :href="`mailto:${emailAddress}`">
            <span>{{ emailAddress }}</span>
          </a>
        </span>
        <span class="phone-number">
          <img src="@/assets/images/icon-phone.svg" />
          <a class="router-link" :href="`tel:${phoneNumber}`">
            <span>{{ phoneNumber }}</span>
          </a>
        </span>
      </div>
      <div class="d-flex align-items-center">
        <router-link class="router-link" to="/login" v-if="!userInfo">
          <span class="button-login">
            <span>{{ $t('app.header.headingBar.login') }}</span>
            <img src="@/assets/images/icon-user.svg" />
          </span>
        </router-link>
        <el-dropdown trigger="click" @command="onClickUser" v-else>
          <span class="button-login">
            <span>{{ `${userInfo.firstName} ${userInfo.lastName}` }}</span>
            <img src="@/assets/images/icon-user.svg" />
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item :command="UserDropdownCommand.PROFILE">{{
                $t('app.header.headingBar.profile')
              }}</el-dropdown-item>
              <el-dropdown-item :command="UserDropdownCommand.ORDER_LIST">{{
                $t('app.header.headingBar.orderList')
              }}</el-dropdown-item>
              <el-dropdown-item :command="UserDropdownCommand.LOGOUT">{{
                $t('app.header.headingBar.logout')
              }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
        <router-link class="router-link" to="/checkout">
          <span class="button-cart">
            <img src="@/assets/images/icon-cart.svg" />
          </span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
import { emailAddress, phoneNumber, UserDropdownCommand } from '../../constants';

export default {
  setup() {
    return { emailAddress, phoneNumber, UserDropdownCommand };
  },
  computed: {
    ...mapGetters('user', ['userInfo']),
  },
  methods: {
    ...mapActions('auth', ['logout']),
    onClickUser(command) {
      switch (command) {
        case UserDropdownCommand.PROFILE:
          this.$router.push('#');
          break;
        case UserDropdownCommand.ORDER_LIST:
          this.$router.push('/order');
          break;
        case UserDropdownCommand.LOGOUT:
          this.logout();
          this.$router.push('/home');
          break;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.container-fluid {
  height: 40px;
  padding-top: 2px;
  color: $color-white;
  font-size: 15px;
  font-weight: 500;
  background-color: $color-primary;
  .heading-bar-content {
    height: 100%;
  }
}
.email-address,
.phone-number {
  margin-right: 20px;
  img {
    margin-right: 5px;
    margin-bottom: 3px;
  }
}
.el-dropdown-menu {
  width: 160px;
}
.button-login {
  margin-right: 20px;
  color: $color-white;
  &:hover {
    cursor: pointer;
    opacity: 0.9;
  }
  img {
    margin-left: 4px;
    margin-bottom: 4px;
  }
}
.button-cart img {
  margin-bottom: 3px;
  &:hover {
    opacity: 0.9;
  }
}
</style>
