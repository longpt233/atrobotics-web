<template>
  <div class="container-fluid bg-primary">
    <div class="container">
      <div class="d-flex justify-content-between p-3">
        <div class="d-flex align-items-center">
          <img src="@/assets/images/icon-email.png" class="me-2" />
          <div class="me-5">tester123@gmail.com</div>
          <img src="@/assets/images/icon-phone.png" class="me-2" />
          <div>0123456789</div>
        </div>

        <div class="d-flex align-items-center justify-content-center">
          <div class="me-2 button button-login">Login</div>
          <img src="@/assets/images/icon-user.png" class="me-4" />
          <div class="me-2 button button-wishlist">Wishlist</div>
          <img src="@/assets/images/icon-favourite.png" class="me-4" />
          <img src="@/assets/images/icon-cart.png" class="button button-cart" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {}
</script>

<style lang="scss" scoped>
* {
  font-size: 16px;
  line-height: 16px;
  color: $color-white;
}

img {
  width: 16px;
  height: 16px;
}

.button {
  cursor: pointer;
}
</style>
