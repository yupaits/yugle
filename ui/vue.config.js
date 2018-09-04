const appName = '鱼钩资讯';

module.exports = {
  pages: {
    index: {
      entry: 'src/index/main.js',
      title: appName
    },
    login: {
      entry: 'src/login/main.js',
      title: '登录 | ' + appName
    },
    404: {
      entry: 'src/404/main.js',
      title: '404 | ' + appName
    }
  },
  baseUrl: "/",
  assetsDir: 'static'
};