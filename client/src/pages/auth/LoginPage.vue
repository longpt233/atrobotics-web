<template>
  <div class="login-page-box">
    <el-card class="login-card" :body-style="{ padding: '60px' }">
      <div class="login-card__title">{{ $t('auth.login.title') }}</div>
      <div class="form d-flex flex-column">
        <el-input
          v-model="email"
          :class="{ 'input-error': errors.email }"
          :placeholder="$t('auth.login.placeholderEmail')"
          type="email"
          size="large"
          @input="onChangeInput()"
        ></el-input>
        <span class="error-message" v-show="errors.email">{{
          capitalize(errors.email)
        }}</span>
        <el-input
          v-model="password"
          :class="{ 'input-error': errors.password }"
          :placeholder="$t('auth.login.placeholderPassword')"
          type="password"
          size="large"
          @input="onChangeInput()"
        ></el-input>
        <span class="error-message" v-show="errors.password">{{
          capitalize(errors.password)
        }}</span>
      </div>
      <div class="login-card__button-forgot-password">
        {{ $t('auth.login.forgotPassword') }}
      </div>
      <el-button class="login-card__button-submit" type="primary" @click="onSubmit">{{
        $t('auth.login.title')
      }}</el-button>
      <div class="login-card__redirect-page">
        <span>{{ `${$t('auth.login.redirectPage.title')} ` }}</span>
        <router-link class="router-link" to="/signup">
          <span class="button-register">{{
            $t('auth.login.redirectPage.labelButton')
          }}</span>
        </router-link>
      </div>
    </el-card>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';
import * as yup from 'yup';
import { useField, useForm } from 'vee-validate';
import { nextTick } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import globalMixin from '@/common/mixins';
import { capitalize } from 'lodash';

export default {
  mixins: [globalMixin],
  setup() {
    const router = useRouter();
    const store = useStore();
    const { t } = useI18n();

    const initialValues = {
      email: '',
      password: '',
    };

    const validations = yup.object({
      email: yup.string().required().email(),
      password: yup.string().required(),
    });

    const { handleSubmit, errors } = useForm({
      validationSchema: validations,
      initialValues: initialValues,
    });

    const onSubmit = handleSubmit(async ({ email, password }) => {
      const error = await store.dispatch('auth/login', {
        email: email,
        password: password,
      });
      if (!error) {
        globalMixin.methods.showSuccessNotification(
          t('auth.login.notification.success.title'),
          t('auth.login.notification.success.message'),
        );
        router.push('/home');
      } else {
        globalMixin.methods.showErrorNotification(
          t('auth.login.notification.fail.title'),
          error,
        );
      }
    });

    const { value: email } = useField('email');
    const { value: password } = useField('password');

    return {
      email,
      password,
      errors,
      onSubmit,
      capitalize,
    };
  },
  computed: {
    ...mapGetters('user', ['userInfo']),
  },
  created() {
    nextTick(() => {
      // redirect to home page if existed a valid token
      if (this.userInfo) {
        this.$router.replace('/');
      }
    });
  },
};
</script>

<style lang="scss" scoped>
.login-page-box {
  width: 450px;
  margin: 180px auto 100px;
  text-align: center;
}

.login-card {
  &__title {
    margin-bottom: 14px;
    font-size: 32px;
    font-weight: 700;
  }
  .form {
    gap: 16px;
    margin-bottom: 16px;
    :deep(.el-input__wrapper) {
      padding: 0;
      border-radius: 5px;
    }
    :deep(.el-input__inner) {
      padding: 0 15px;
      border-radius: 5px;
    }
  }
  &__button-forgot-password {
    margin-bottom: 10px;
    font-weight: 500;
    color: $color-blue;
    &:hover {
      text-decoration: underline;
      cursor: pointer;
    }
  }
  &__button-submit {
    margin-bottom: 14px;
    padding: 18px 15px;
    font-size: 18px;
    border-radius: 5px;
  }
  &__redirect-page {
    .button-register {
      font-weight: 500;
      color: $color-blue;
      &:hover {
        text-decoration: underline;
        cursor: pointer;
      }
    }
  }
}
</style>
