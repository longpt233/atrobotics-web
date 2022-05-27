<template>
  <div class="register-page-box">
    <el-card class="register-card" :body-style="{ padding: '50px' }">
      <div class="register-card__title">{{ $t('auth.register.title') }}</div>
      <div class="register-card__subtitle">{{ $t('auth.register.subtitle') }}</div>
      <div class="form d-flex flex-column">
        <div class="input-name d-flex">
          <el-input
            v-model="firstName"
            :placeholder="$t('auth.register.placeholderFirstName')"
            type="text"
            size="large"
          />
          <el-input
            v-model="lastName"
            :placeholder="$t('auth.register.placeholderLastName')"
            type="text"
            size="large"
          />
        </div>
        <span class="error-message" v-show="errors.firstName">
          {{ capitalize(errors.firstName) }}
        </span>
        <span class="error-message" v-show="errors.lastName">
          {{ capitalize(errors.lastName) }}</span
        >
        <el-input
          v-model="phone"
          :placeholder="$t('auth.register.placeholderPhone')"
          type="phone"
          size="large"
        />
        <span class="error-message" v-show="errors.phone">{{
          capitalize(errors.phone)
        }}</span>
        <el-input
          v-model="email"
          :placeholder="$t('auth.register.placeholderEmail')"
          type="email"
          size="large"
          autocomplete="new-password"
        />
        <span class="error-message" v-show="errors.email">{{
          capitalize(errors.email)
        }}</span>
        <el-input
          v-model="password"
          :placeholder="$t('auth.register.placeholderPassword')"
          type="password"
          size="large"
          autocomplete="new-password"
        />
        <span class="error-message" v-show="errors.password">{{
          capitalize(errors.password)
        }}</span>
        <el-input
          v-model="repeatPassword"
          :placeholder="$t('auth.register.placeholderRepeatPassword')"
          type="password"
          size="large"
          autocomplete="new-password"
        />
        <span class="error-message" v-show="errors.repeatPassword">{{
          capitalize(errors.repeatPassword)
        }}</span>
      </div>
      <el-button class="register-card__button-submit" type="primary" @click="onSubmit">{{
        $t('auth.register.title')
      }}</el-button>
      <div class="register-card__redirect-page">
        <span>{{ `${$t('auth.register.redirectPage.title')} ` }}</span>
        <router-link class="router-link" to="/login">
          <span class="button-register">{{
            $t('auth.register.redirectPage.labelButton')
          }}</span>
        </router-link>
      </div>
    </el-card>
  </div>
</template>

<script>
import globalMixin from '@/common/mixins';
import { useI18n } from 'vue-i18n';
import * as yup from 'yup';
import { useField, useForm } from 'vee-validate';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';
import { capitalize } from 'lodash';

import { FORM_VALIDATION } from '@/common/constants';

export default {
  mixins: [globalMixin],
  setup() {
    const router = useRouter();
    const store = useStore();
    const { t } = useI18n();

    const initialValues = {
      firstName: '',
      lastName: '',
      phone: '',
      email: '',
      password: '',
      repeatPassword: '',
    };

    const validations = yup.object({
      firstName: yup.string().required().trim(),
      lastName: yup.string().required().trim(),
      phone: yup.string().required().trim(),
      email: yup.string().required().email().trim(),
      password: yup.string().required().min(FORM_VALIDATION.passwordMinLength).trim(),
      repeatPassword: yup
        .string()
        .oneOf([yup.ref('password'), null], 'Passwords must match'),
    });

    const { handleSubmit, errors } = useForm({
      validationSchema: validations,
      initialValues: initialValues,
    });

    const onSubmit = handleSubmit(
      async ({ firstName, lastName, phone, email, password }) => {
        const error = await store.dispatch('auth/register', {
          firstname: firstName,
          lastname: lastName,
          phone: phone,
          email: email,
          password: password,
        });
        if (!error) {
          globalMixin.methods.showSuccessNotification(
            t('auth.register.notification.success.title'),
            t('auth.register.notification.success.message'),
          );
          router.push('/login');
        } else {
          globalMixin.methods.showErrorNotification(
            t('auth.register.notification.fail.title'),
            error,
          );
        }
      },
    );
    const { value: firstName } = useField('firstName');
    const { value: lastName } = useField('lastName');
    const { value: phone } = useField('phone');
    const { value: email } = useField('email');
    const { value: password } = useField('password');
    const { value: repeatPassword } = useField('repeatPassword');

    return {
      firstName,
      lastName,
      phone,
      email,
      password,
      repeatPassword,
      errors,
      onSubmit,
      capitalize,
    };
  },
};
</script>

<style lang="scss" scoped>
.register-page-box {
  width: 520px;
  margin: 140px auto 50px;
  text-align: center;
}
.register-card {
  &__title {
    font-size: 32px;
    font-weight: 700;
  }
  &__subtitle {
    margin-bottom: 15px;
    font-weight: 400;
    color: $color-dark-grey;
  }
  .form {
    gap: 15px;
    margin-bottom: 22px;
    .input-name {
      gap: 15px;
    }
    :deep(.el-input__wrapper) {
      padding: 0;
      border-radius: 5px;
    }
    :deep(.el-input__inner) {
      padding: 0 15px;
      border-radius: 5px;
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
