<template>
  <div>
    <h2>系统设置</h2>
    <a-tabs type="card" tabPosition="left">
      <a-tab-pane v-for="item in $consts.settingsMenu" :tab="item.label" :key="item.key">
        <user-manage v-if="item.key === 'user'" key="user"></user-manage>
        <role-manage v-else-if="item.key === 'role'" key="role"></role-manage>
        <permission-manage v-else-if="item.key === 'permission'" key="permission"></permission-manage>
        <modify-password v-else-if="item.key === 'password'" key="password"></modify-password>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
  import UserManage from "./settings/UserManage";
  import RoleManage from "./settings/RoleManage";
  import PermissionManage from "./settings/PermissionManage";
  import ModifyPassword from "./settings/ModifyPassword";
  export default {
    name: "Settings",
    components: {UserManage, RoleManage, PermissionManage, ModifyPassword},
  }
</script>

<style scoped>

</style>