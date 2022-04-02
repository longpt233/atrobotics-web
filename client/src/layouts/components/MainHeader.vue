<template>
  <div class="container-fluid bg-primary">
    <div class="container">
      <div class="d-flex justify-content-between p-3">
        <div class="d-flex align-items-center">
          <img src="@/assets/images/icon-email.png" class="me-2 mb-1" />
          <div class="me-5"><strong>tester123@gmail.com</strong></div>
          <img src="@/assets/images/icon-phone.png" class="me-2 mb-1" />
          <div><strong>0123456789</strong></div>
        </div>

        <div class="d-flex align-items-center justify-content-center">
          <div class="button button-login me-2"><strong>Login</strong></div>
          <img src="@/assets/images/icon-user.png" class="me-4 mb-1" />
          <div class="button button-wishlist me-2"><strong>Wishlist</strong></div>
          <img src="@/assets/images/icon-favourite.png" class="me-4 mb-1" />
          <img src="@/assets/images/icon-cart.png" class="button button-wishlist mb-1" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
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
