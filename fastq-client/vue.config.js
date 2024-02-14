const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: ["vuetify"],
  devServer: {
    proxy: {
      '/server': {
        target: 'http://localhost:8090/', // provide proxy  for your project
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/server': ''
        }
      },
      '/local': {
        target: 'http://localhost:3000/', // provide proxy  for your project
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/local': ''
        }
      },
    }
  },
});
