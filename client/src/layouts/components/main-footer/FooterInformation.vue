<template>
  <div class="footer-information d-flex flex-column">
    <div class="footer-information__title">
      {{ $t('app.footer.information.title') }}
    </div>
    <div class="footer-information__phone d-flex align-items-center">
      <el-icon color="#0d6efd"><PhoneIcon /></el-icon>
      <a class="router-link" :href="`tel:${phoneNumber}`">
        <span>{{ phoneNumber }}</span>
      </a>
    </div>
    <div class="footer-information__email d-flex align-items-center">
      <el-icon color="#0d6efd"><MessageIcon /></el-icon>
      <a class="router-link" :href="`mailto:${emailAddress}`">
        <span>{{ emailAddress }}</span>
      </a>
    </div>
    <div class="footer-information__address d-flex align-items-center">
      <el-icon color="#0d6efd"><LocationIcon /></el-icon>
      <a class="router-link" :href="googleMapUrl">
        <span>{{ $t('app.footer.information.address') }}</span>
      </a>
    </div>
  </div>
</template>

<script>
import { emailAddress, phoneNumber, googleMapUrl } from '../../constants';
import {
  Phone as PhoneIcon,
  Message as MessageIcon,
  Location as LocationIcon,
} from '@element-plus/icons-vue';

export default {
  setup() {
    return { emailAddress, phoneNumber, googleMapUrl };
  },
  components: { PhoneIcon, MessageIcon, LocationIcon },
};
</script>

<style lang="scss" scoped>
.footer-information {
  gap: 10px;
  font-size: 14px;
  &__title {
    color: $color-primary;
    font-size: 20px;
    font-weight: 700;
  }
  i {
    margin-right: 10px;
  }
}
</style>
