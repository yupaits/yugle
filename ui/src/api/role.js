import request from '../request'

const rolePathPrefix = "/role";

const role = {
  getRolePage(page, size, keyword) {
    return request.api.get(`${rolePathPrefix}/page?page=${page}&size=${size}&keyword=${keyword}`);
  },
  listRoles() {
    return request.api.get(`/roles`);
  },
  addRole(role) {
    return request.api.post(`${rolePathPrefix}`, role);
  },
  updateRole(roleId, role) {
    return request.api.put(`${rolePathPrefix}/${roleId}`, role);
  },
  getRolePermissions(roleId) {
    return request.api.get(`${rolePathPrefix}_permissions/${roleId}`);
  },
  assignPermissions(roleId, permissionIds) {
    return request.api.post(`${rolePathPrefix}/permissions/assign?roleId=${roleId}`, permissionIds);
  },
  deleteRole(roleId) {
    return request.api.delete(`${rolePathPrefix}/${roleId}`);
  }
};

export default role