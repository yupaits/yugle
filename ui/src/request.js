import {notification} from 'ant-design-vue'
import axios from 'axios'
import consts from './consts'
import 'vue-cookies'

import 'ant-design-vue/lib/notification/style/index.css'
import 'ant-design-vue/lib/icon/style/css'

//认证授权请求
const auth = axios.create({
  baseURL: consts.urls.authBaseUrl
});

auth.interceptors.response.use(res => {
  if (res.status !== 200) {
    return Promise.reject(res);
  }
  if (res.data.code === 200) {
    //返回自定义成功Result，示例：{code: 200, message: '成功', data: []}
    return Promise.resolve(res.data);
  }
  if (res.data.code === undefined) {
    //返回默认response
    return Promise.resolve(res);
  }
  //返回自定义失败Result
  notification.error({
    message: '错误',
    description: res.data.msg,
    duration: 3
  });
  return Promise.reject(res);
}, error => {
  if (!error.response) {
    //无状态码的错误信息
    notification.error({
      message: '错误',
      description: error.toString(),
      duration: 3
    });
    return Promise.reject(error.response);
  }
  //带状态码的错误信息
  if (error.response.data) {
    //带返回内容的错误信息
    notification.error({
      message: '错误',
      description: error.response.data.msg,
      duration: 3
    });
    return Promise.reject(error.response.data);
  }
  notification.error({
    message: '错误',
    description: error.response.status + ' - ' + error.response.statusText,
    duration: 3
  });
  return Promise.reject(error.response);
});

//刷新AccessToken
function refreshToken() {
  auth.post('/refresh_token', undefined, {
    headers: {
      Authorization: consts.token.prefix + window.$cookies.get(consts.token.cookieName)
    }
  }).then(res => {
    window.$cookies.set(consts.token.cookieName, res.token, consts.token.maxRefreshTime);
    window.$cookies.set(consts.token.refreshFlag, 'flag', consts.token.expiredTime);
  });
}

//API请求
const api = axios.create({
  baseURL: consts.urls.apiBaseUrl
});

api.interceptors.request.use(config => {
  //查看是否存在刷新标记，如果不存在则需要刷新token
  const refreshFlag = window.$cookies.get(consts.token.refreshFlag);
  alert(refreshFlag);
  if (refreshFlag === null) {
    refreshToken();
  }
  config.headers.Authorization = consts.token.prefix + window.$cookies.get(consts.token.cookieName);
  return config;
}, error => {
  return Promise.reject(error);
});

api.interceptors.response.use(res => {
  if (res.status !== 200) {
    return Promise.reject(res);
  }
  if (res.data.code === 200) {
    //返回自定义成功Result，示例：{code: 200, message: '成功', data: []}
    return Promise.resolve(res.data);
  }
  if (res.data.code === undefined) {
    //返回默认response
    return Promise.resolve(res);
  }
  //返回自定义失败Result
  notification.error({
    message: '错误',
    description: res.data.msg,
    duration: 3
  });
  return Promise.reject(res);
}, error => {
  if (!error.response) {
    //无状态码的错误信息
    notification.error({
      message: '错误',
      description: error.toString(),
      duration: 3
    });
    return Promise.reject(error.response);
  }
  //带状态码的错误信息
  if (error.response.status === 401) {
    window.location.href = '/login';
  }
  if (error.response.data) {
    //带返回内容的错误信息
    notification.error({
      message: '错误',
      description: error.response.data.code + ' - ' + error.response.data.msg,
      duration: 3
    });
    return Promise.reject(error.response.data);
  }
  notification.error({
    message: '错误',
    description: error.response.status + ' - ' + error.response.statusText,
    duration: 3
  });
  return Promise.reject(error.response);
});

export default {
  auth, api
}