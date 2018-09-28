<template>
  <div>
    <a-row class="opt-container">
      <a-button class="btn-success" icon="plus" @click="addRole">添加角色</a-button>
      <div class="pull-right">
        <span>关键字：
          <a-input v-model="keyword" placeholder="请输入关键字" class="search-input mr-2"></a-input>
        </span>
        <a-button icon="search" @click="handleSearch">搜索</a-button>
      </div>
    </a-row>
    <a-table size="small"
             :columns="columns"
             :rowKey="record => record.ID"
             :dataSource="roles"
             :pagination="false"
             :loading="loading"
             :scroll="$consts.tableScroll.page">
      <template slot="created-time" slot-scope="record">
        {{record.CreatedAt | goTimeFormat}}
      </template>
      <template slot="operation" slot-scope="record">
        <a-button size="small" class="mr-1" @click="editRole(record)">编辑</a-button>
        <a-popconfirm title="确定要删除该角色吗？" placement="topRight"
                      @confirm="deleteRole(record.ID)" okText="确定" cancelText="取消">
          <a-button type="danger" size="small" class="mr-1">删除</a-button>
        </a-popconfirm>
        <a-button size="small" @click="assignPermissions(record.ID)">分配权限</a-button>
      </template>
    </a-table>
    <a-pagination size="small"
                  class="pager"
                  v-model="pager.current"
                  :total="pager.total"
                  :pageSize.sync="pager.pageSize"
                  :pageSizeOptions="pager.pageSizeOptions"
                  :showQuickJumper="pager.showQuickJumper"
                  :showSizeChanger="pager.showSizeChanger"
                  @change="handlePagerChange"></a-pagination>

    <a-modal :title="roleModal.title"
             :visible="modalVisible"
             :maskClosable="false"
             @ok="roleModal.ok"
             @cancel="() => {this.modalVisible = false}">
      <a-form>
        <a-form-item label="Key" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol">
          <a-input v-model="role.key" placeholder="请填写角色Key，一个角色对应的Key是唯一的"></a-input>
        </a-form-item>
        <a-form-item label="角色名称" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol">
          <a-input v-model="role.name" placeholder="请填写角色名称"></a-input>
        </a-form-item>
        <a-form-item label="角色描述" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol">
          <a-textarea v-model="role.description" :autosize="{minRows: 3, maxRows: 5}" placeholder="请填写该角色的描述信息"></a-textarea>
        </a-form-item>
      </a-form>
    </a-modal>

    <a-drawer width="480" placement="right"
              title="分配权限"
              class="drawer-container"
              :closable="false"
              @close="() => {this.permissionsVisible = false}"
              :visible="permissionsVisible">
      <a-checkbox-group @change="handleCheckPermissions">
        <a-checkbox v-for="permission in allPermissions" :key="permission.Key" :value="permission.ID">
          <span>{{permission.Name}}</span>
          <a-tag>{{permission.PermType}}</a-tag>
        </a-checkbox>
      </a-checkbox-group>
      <div class="drawer-opt">
        <a-button class="mr-1" @click="() => {this.permissionsVisible = false}">取消</a-button>
        <a-button type="primary" @click="handleAssignPermissions">确认</a-button>
      </div>
    </a-drawer>
  </div>
</template>

<script>
  const columns = [
    {title: '角色Key', width: '15%', dataIndex: 'Key'},
    {title: '角色名', width: '15%', dataIndex: 'Name'},
    {title: '角色描述', width: '40%', dataIndex: 'Description'},
    {title: '创建时间', width: '15%', scopedSlots: {customRender: 'created-time'}},
    {title: '操作', width: '20%', scopedSlots: {customRender: 'operation'}},
  ];
  export default {
    name: "RoleManage",
    data() {
      return {
        columns: columns,
        roles: [],
        allPermissions: [],
        rolePermissions: [],
        pager: {
          showQuickJumper: true,
          showSizeChanger: true,
          current: 1,
          total: 1,
          pageSize: parseInt(this.$consts.pageSizes[0]),
          pageSizeOptions: this.$consts.pageSizes
        },
        loading: false,
        keyword: '',
        roleModal: {
          title: '',
          ok: () => {}
        },
        modalVisible: false,
        permissionsVisible: false,
        role: {}
      }
    },
    mounted() {
      this.fetchRoles();
      this.fetchPermissions();
    },
    methods: {
      fetchRoles() {
        this.loading = true;
        this.$api.role.getRolePage(this.pager.current, this.pager.pageSize, this.keyword).then(res => {
          this.roles = res.data.Content;
          this.pager.current = res.data.Page;
          this.pager.total = res.data.Total;
          this.loading = false;
        }).catch(() => {
          this.loading = false;
        });
      },
      fetchPermissions() {
        this.$api.permission.listPermissions().then(res => {
          this.allPermissions = res.data;
        });
      },
      handleSearch() {
        this.fetchRoles();
      },
      addRole() {
        this.role = {};
        this.roleModal = {
          title: '添加角色',
          ok: this.handleAddRole
        };
        this.modalVisible = true;
      },
      editRole(role) {
        this.role = {
          id: role.ID,
          key: role.Key,
          name: role.Name,
          description: role.Description
        };
        this.roleModal = {
          title: '编辑角色',
          ok: this.handleEditRole
        };
        this.modalVisible = true;
      },
      deleteRole(roleId) {
        this.$api.role.deleteRole(roleId).then(() => {
          this.$message.success('删除角色成功');
        });
      },
      assignPermissions(roleId) {
        this.$api.role.getRolePermissions(roleId).then(res => {
          this.rolePermissions = res.data;
        });
        this.permissionsVisible = true;
      },
      handleCheckPermissions(permissionIds) {
        this.rolePermissions = permissionIds;
      },
      handleAddRole() {
        this.$api.role.addRole(this.role).then(() => {
          this.$message.success('添加角色成功');
          this.modalVisible = false;
          this.fetchRoles();
        });
      },
      handleEditRole() {
        this.$api.role.updateRole(this.role.id, this.role).then(() => {
          this.$message.success('更新角色成功');
          this.modalVisible = false;
          this.fetchRoles();
        });
      },
      handleAssignPermissions() {

      },
      handlePagerChange(page, pageSize) {
        this.pager.current = page;
        this.pager.pageSize = pageSize;
        this.fetchRoles();
      }
    }
  }
</script>

<style scoped>

</style>