import request from '../request'

const permPathPrefix = "/permission";

const permission = {
  getPermissionPage(page, size, permissionQuery) {
    return request.api.post(`${permPathPrefix}/page?page=${page}&size=${size}`, permissionQuery);
  },
  listPermissions() {
    return request.api.get(`/permissions`);
  },
  addPermission(permission) {
    return request.api.post(`${permPathPrefix}`, permission);
  },
  updatePermission(permissionId, permission) {
    return request.api.put(`${permPathPrefix}/${permissionId}`, permission);
  },
  deletePermission(permissionId) {
    return request.api.delete(`${permPathPrefix}/${permissionId}`);
  }
};

export default permission