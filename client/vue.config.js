const { defineConfig } = require('@vue/cli-service');
module.exports = defineConfig({
  transpileDependencies: true,
  css: {
    loaderOptions: {
      sass: {
        additionalData: `
            @use "@/assets/styles/variables.scss" as *;
            @use "@/assets/styles/mixins.scss" as *;
            `,
      },
    },
  },
});
