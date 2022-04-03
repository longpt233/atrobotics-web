<template>
  <div class="container-fluid bg-primary">
    <div class="container d-flex justify-content-between p-2">
      <div class="d-flex align-items-center">
        <div class="ms-1 me-5">
          <img class="me-2 mb-1" src="@/assets/images/icon-email.svg" />
          <strong>tester123@gmail.com</strong>
        </div>
        <div>
          <img class="me-2 mb-1" src="@/assets/images/icon-phone.svg" />
          <strong>0123456789</strong>
        </div>
      </div>
      <div class="d-flex align-items-center">
        <div class="button">
          <strong>Đăng nhập</strong>
          <img class="ms-2 me-4 mb-1" src="@/assets/images/icon-user.svg" />
        </div>
        <img
          class="button mb-1 me-2"
          src="@/assets/images/icon-cart.svg"
          style="width: 24px; height: auto"
        />
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
* {
  font-size: 18px;
  line-height: 16px;
  color: $color-white;
}

.button {
  img {
    width: 16px;
    height: auto;
  }
}
</style>
