import dateFns from 'date-fns'

const urls = {
  authBaseUrl: '/auth',
  apiBaseUrl: '/api'
};

const token = {
  prefix: 'Bearer ',
  cookieName: 'AccessToken',
  needRefresh: 'NeedRefreshFlag',
  expiredTime: '10s',
  maxRefreshTime: '30s'
};

const form = {
  labelCol: {span: 3},
  wrapperCol: {span: 15},
  submitCol: {span: 15, push: 3}
};

const fluidForm = {
  labelCol: {span: 4},
  wrapperCol: {span: 20},
  submitCol: {span: 20, push: 4}
};

const tableScroll = {
  page: {y: 560},
  card: {y: 300}
};

const navs = [
  {
    label: '美图', icon: 'picture', key: 'picture', items: [
      {label: '必应壁纸', to: '/picture/bing', key: 'bing-picture'},
      {label: 'Shot on OnePlus', to: '/picture/shot_on_oneplus', key: 'shot-on-oneplus'},
    ]
  },
];

const userMenu = [
  {label: '个人中心', icon: 'idcard', to: '/profile'},
  {label: '系统设置', icon: 'setting', to: '/settings'},
];

const settingsMenu = [
  {label: '用户管理', key: 'user'},
  {label: '角色管理', key: 'role'},
  {label: '权限管理', key: 'permission'},
  {label: '修改密码', key: 'password'},
];

const pageSizes = ['10', '20', '50', '100'];

const golangDateTimeFormat = function (date) {
  if (date) {
    if (dateFns.getYear(date) === 1) {
      return '';
    }
    return dateFns.format(date, 'YYYY-MM-DD HH:mm:ss');
  }
  return '';
};

export default {
  urls, token, form, fluidForm, tableScroll, navs, userMenu, settingsMenu, pageSizes,
  golangDateTimeFormat
}