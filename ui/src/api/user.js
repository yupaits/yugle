import request from '../request'

const userPathPrefix = '/user';

const user = {
  getUserPage(page, size, userQuery) {
    return request.api.post(`${userPathPrefix}/page?page=${page}&size=${size}`, userQuery);
  },
  getCurrentUser() {
    return request.api.get(`${userPathPrefix}/current`);
  },
  updateUser(userId, user) {
    return request.api.put(`${userPathPrefix}/${userId}`, user);
  },
  changeUserStatus(userId, status) {
    return request.api.put(`${userPathPrefix}/${userId}/status?enabled=${status}`);
  },

};

export default user