<template>
  <div class="product-images-box">
    <Splide :options="splideOptions">
      <SplideSlide v-for="imageUrl in productDetail?.images" :key="imageUrl">
        <img :src="imageUrl" />
      </SplideSlide>
    </Splide>
  </div>
</template>

<script>
import { Splide, SplideSlide } from '@splidejs/vue-splide';
import { mapGetters } from 'vuex';

export default {
  components: { Splide, SplideSlide },
  setup() {
    const splideOptions = { rewind: true, autoplay: true, interval: 1000, arrows: false };
    return { splideOptions };
  },
  computed: {
    ...mapGetters('productDetail', ['productDetail']),
  },
};
</script>

<style lang="scss" scoped>
.product-images-box {
  flex-shrink: 0;
  width: 500px;
  text-align: center;
  img {
    max-width: 100%;
    height: 400px;
  }
}
</style>
