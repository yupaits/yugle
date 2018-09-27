import request from '../request'

const userPathPrefix = '/user';
const rolePathPrefix = "/role";

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
  getRolePage(page, size, keyword) {
    return request.api.get(`${rolePathPrefix}/page?page=${page}&size=${size}&keyword=${keyword}`);
  },
  listRoles() {
    return request.api.get(`/roles`);
  }
};

export default user