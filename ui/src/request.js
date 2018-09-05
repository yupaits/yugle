import {notification} from 'ant-design-vue'
import axios from 'axios'
import consts from './consts'

const auth = axios.create({
  baseURL: consts.urls.authBaseUrl
});

const api = axios.create({
  baseURL: consts.urls.apiBaseUrl
});

export default {
  auth, api
}