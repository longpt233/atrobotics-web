import { ElNotification } from 'element-plus';

export default {
  methods: {
    parsePrice(price) {
      return new Intl.NumberFormat('it-IT', {
        style: 'decimal',
        currency: 'VND',
      }).format(price);
    },
    showSuccessNotification(title, message) {
      ElNotification({
        type: 'success',
        title: title,
        message: message,
      });
    },
    showWarningNotification(title, message) {
      ElNotification({
        type: 'warning',
        title: title,
        message: message,
      });
    },
    showErrorNotification(title, message) {
      ElNotification({
        type: 'error',
        title: title,
        message: message,
      });
    },
    translateYupError(yupError) {
      if (!yupError) return '';
      if (yupError?.i18nKey)
        return this.$t(yupError?.i18nKey, {
          ...yupError?.params,
        });
      return this.$t(yupError);
    },
  },
};
