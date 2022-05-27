<template>
  <el-autocomplete
    ref="searchBox"
    v-model="state"
    :fetch-suggestions="onChangeKeyword"
    :placeholder="placeholder"
    @select="onSelectKeyword"
    @keyup.enter="onSelectKeyword"
  >
    <template #append v-if="isShowSearchButton">
      <el-button type="primary" @click="onSelectKeyword">
        <el-icon> <IconSearch /> </el-icon>
      </el-button> </template
  ></el-autocomplete>
</template>

<script>
import { Search as IconSearch } from '@element-plus/icons-vue';

export default {
  components: { IconSearch },
  props: {
    placeholder: String,
    isShowSearchButton: {
      type: Boolean,
      default: true,
    },
  },
  data() {
    return {
      state: '',
    };
  },
  methods: {
    onChangeKeyword(keyword, callback) {
      if (keyword.trim() !== '') {
        this.$emit('on-change-keyword', { keyword: keyword.trim(), callback });
      } else {
        callback([]);
      }
    },
    onSelectKeyword() {
      this.$refs.searchBox.suggestions = [];
      this.$emit('on-select-keyword', this.state);
      this.state = '';
    },
  },
};
</script>

<style lang="scss" scoped>
button {
  background-color: $color-primary !important;
  color: $color-white !important;
}
</style>
