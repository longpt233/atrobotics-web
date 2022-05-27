<template>
  <div class="container">
    <Splide :options="splideOptions">
      <SplideSlide v-for="banner in bannerList" :key="banner.id">
        <router-link class="router-link" :to="`/product/${banner.productId}`">
          <img :src="banner.image" />
        </router-link>
      </SplideSlide>
    </Splide>
  </div>
</template>

<script>
import { Splide, SplideSlide } from '@splidejs/vue-splide';
import { mapGetters } from 'vuex';

export default {
  components: {
    Splide,
    SplideSlide,
  },
  computed: {
    ...mapGetters('home', ['bannerList']),
  },
  setup() {
    const splideOptions = { rewind: true, autoplay: true, interval: 2000, arrows: false };
    return { splideOptions };
  },
};
</script>

<style lang="scss" scoped>
.container {
  margin-top: 105px;
  margin-bottom: 20px;
  text-align: center;
  img {
    width: auto;
    height: 450px;
  }
}
</style>
