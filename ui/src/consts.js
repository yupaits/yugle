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

export default {
  urls, token, form, fluidForm
}