<template>
  <router-view />
</template>

<script>
import { mapActions } from 'vuex';

export default {
  created() {
    this.autoLogin();
  },
  methods: {
    ...mapActions('auth', { actionAutoLogin: 'autoLogin' }),
    async autoLogin() {
      const token = localStorage.getItem('token');
      if (!token || token.trim() === '') {
        return;
      }
      try {
        await this.actionAutoLogin(token);
      } catch (error) {
        console.log(error);
      }
    },
  },
};
</script>

<style lang="scss"></style>
