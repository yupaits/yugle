const urls = {
  authBaseUrl: '/auth',
  apiBaseUrl: '/api'
};

const token = {
  prefix: 'Bearer ',
  cookieName: 'AccessToken',
  refreshFlag: 'TokenRefreshFlag',
  expiredTime: '1h',
  maxRefreshTime: '48h'
};

const form = {
  labelCol: {span: 3},
  wrapperCol: {span: 15},
  submitCol: {span: 15, push: 3}
};

const fluidForm = {
  labelCol: {span: 3},
  wrapperCol: {span: 21},
  submitCol: {span: 21, push: 3}
};

const navs = [
  {
    label: '美图', icon: 'icon-meitu', key: 'picture', items: [
      {label: '必应壁纸', to: '/picture/bing', key: 'bing-picture'},
      {label: 'Shot on OnePlus', to: '/picture/shot_on_oneplus', key: 'shot-on-oneplus'},
    ]
  },
  {
    label: '视频', icon: 'icon-shipin', key: 'video', items: [
      {label: '', to: '/video', key: '-video'}
    ]
  },
  {
    label: '电影', icon: 'icon-dy-light-copy', key: 'movie', items: []
  },
];

const userMenu = [
  {label: '个人中心', icon: 'idcard', to: '/profile'},
  {label: '系统设置', icon: 'setting', to: '/settings'},
];

export default {
  urls, token, form, fluidForm, navs, userMenu
}