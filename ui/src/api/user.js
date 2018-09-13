import request from '../request'

const userPathPrefix = '/user';

const user = {
  getUserPage(page, size, userQuery) {
    return request.api.post(`${userPathPrefix}/page?page=${page}&size=${size}`, userQuery);
  }
};

export default user