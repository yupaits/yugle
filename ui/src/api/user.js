import request from '../request'

const userPathPrefix = '/user';

const user = {
  getUserPage(page, size, userQuery) {
    return request.api.post(`${userPathPrefix}/page?page=${page}&size=${size}`, userQuery);
  },
  getCurrentUser() {
    return request.api.get(`${userPathPrefix}/current`);
  },
  addUser(user) {
    return request.api.post(`${userPathPrefix}`, user);
  },
  updateUser(userId, user) {
    return request.api.put(`${userPathPrefix}/${userId}`, user);
  },
  changeUserStatus(userId, status) {
    return request.api.put(`${userPathPrefix}/${userId}/status?enabled=${status}`);
  },
  modifyPassword(passwordModify) {
    return request.api.post(`${userPathPrefix}/password`, passwordModify);
  },
  getUserRoles(userId) {
    return request.api.get(`${userPathPrefix}_roles/${userId}`);
  },
  assignRoles(userId, roleIds) {
    return request.api.post(`${userPathPrefix}/roles/assign?userId=${userId}`, roleIds);
  }
};

export default user