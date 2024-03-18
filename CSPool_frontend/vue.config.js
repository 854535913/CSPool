const { defineConfig } = require('@vue/cli-service');
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080', //后端服务地址
        changeOrigin: true,
        pathRewrite: { '^/api': '' },
      },
    },
  },
});
