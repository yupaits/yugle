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
    403: {
      entry: 'src/403/main.js',
      title: '403 | ' + appName
    },
    404: {
      entry: 'src/404/main.js',
      title: '404 | ' + appName
    },
    500: {
      entry: 'src/500/main.js',
      title: '500 | ' + appName
    }
  },
  baseUrl: "/",
  assetsDir: 'static'
};