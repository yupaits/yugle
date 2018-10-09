<template>
  <div>
    <a-row class="opt-container">
      <a-button class="btn-success" icon="plus" @click="addPermission">添加权限</a-button>
      <div class="pull-right">
        <span>关键字：
          <a-input v-model="permissionQuery.keyword" placeholder="请输入关键字" class="search-input mr-2"></a-input>
        </span>
        <span>权限类型：
          <a-select v-model="permissionQuery.permType" class="mr-2" @change="handleSearch">
            <a-select-option :value="0">操作</a-select-option>
            <a-select-option :value="1">菜单</a-select-option>
            <a-select-option :value="2">按钮</a-select-option>
          </a-select>
        </span>
        <a-button icon="search" @click="handleSearch">搜索</a-button>
      </div>
    </a-row>
    <a-table size="small"
             :columns="columns"
             :rowKey="record => record.ID"
             :dataSource="permissions"
             :pagination="false"
             :loading="loading"
             :scroll="$consts.tableScroll.page">
      <template slot="perm-type" slot-scope="record">
        <span v-html="permTypeFormat(record.PermType)"></span>
      </template>
      <template slot="created-time" slot-scope="record">
        {{record.CreatedAt | goTimeFormat}}
      </template>
      <template slot="operation" slot-scope="record">
        <a-button size="small" class="mr-1" @click="editPermission(record)">编辑</a-button>
        <a-popconfirm title="确定要删除该权限吗？" placement="topRight"
                      @confirm="deletePermission(record.ID)" okText="确定" cancelText="取消">
          <a-button type="danger" size="small">删除</a-button>
        </a-popconfirm>
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

    <a-modal :title="permissionModal.title"
             :visible="modalVisible"
             :maskClosable="false"
             @ok="permissionModal.ok"
             @cancel="() => {this.modalVisible = false}">
      <a-form>
        <a-form-item label="Key" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol" required>
          <a-input v-model="permission.key" placeholder="请填写权限Key，一个权限对应的Key是唯一的"></a-input>
        </a-form-item>
        <a-form-item label="权限名称" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol" required>
          <a-input v-model="permission.name" placeholder="请填写权限名称"></a-input>
        </a-form-item>
        <a-form-item label="权限类型" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol" required>
          <a-select v-model="permission.permType" placeholder="请选择权限类型" class="form-select">
            <a-select-option :value="0">操作</a-select-option>
            <a-select-option :value="1">菜单</a-select-option>
            <a-select-option :value="2">按钮</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="权限描述" :labelCol="$consts.fluidForm.labelCol" :wrapperCol="$consts.fluidForm.wrapperCol">
          <a-textarea v-model="permission.description" :autosize="{minRows: 3, maxRows: 5}" placeholder="请填写权限描述"></a-textarea>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script>
  const columns = [
    {title: '权限Key', width: '15%', dataIndex: 'Key'},
    {title: '权限名称', width: '15%', dataIndex: 'Name'},
    {title: '权限类型', width: '10%', scopedSlots: {customRender: 'perm-type'}},
    {title: '权限描述', width: '30%', dataIndex: 'Description'},
    {title: '创建时间', width: '15%', scopedSlots: {customRender: 'created-time'}},
    {title: '操作', width: '15%%', scopedSlots: {customRender: 'operation'}},
  ];
  export default {
    name: "PermissionManage",
    data() {
      return {
        columns: columns,
        permissions: [],
        pager: {
          showQuickJumper: true,
          showSizeChanger: true,
          current: 1,
          total: 1,
          pageSize: parseInt(this.$consts.pageSizes[0]),
          pageSizeOptions: this.$consts.pageSizes
        },
        loading: false,
        permissionQuery: {
          keyword: '',
          permType: 0
        },
        permissionModal: {
          title: '',
          ok: () => {}
        },
        modalVisible: false,
        permission: {}
      }
    },
    mounted() {
      this.fetchPermissions();
    },
    methods: {
      fetchPermissions() {
        this.loading = true;
        this.$api.permission.getPermissionPage(this.pager.current, this.pager.pageSize, this.permissionQuery).then(res => {
          this.permissions = res.data.Content;
          this.pager.current = res.data.Page;
          this.pager.total = res.data.Total;
          this.loading = false;
        }).catch(() => {
          this.loading = false;
        });
      },
      handleSearch() {
        this.fetchPermissions();
      },
      addPermission() {
        this.permission = {
          permType: 0
        };
        this.permissionModal = {
          title: '添加权限',
          ok: this.handleAddPermission
        };
        this.modalVisible = true;
      },
      editPermission(permission) {
        this.permission = {
          id: permission.ID,
          key: permission.Key,
          name: permission.Name,
          permType: permission.PermType,
          description: permission.Description
        };
        this.permissionModal = {
          title: '编辑权限',
          ok: this.handleEditPermission
        };
        this.modalVisible = true;
      },
      deletePermission(permissionId) {
        this.$api.permission.deletePermission(permissionId).then(() => {
          this.$message.success('删除权限成功');
          this.fetchPermissions();
        });
      },
      handleAddPermission() {
        this.$api.permission.addPermission(this.permission).then(() => {
          this.$message.success('添加权限成功');
          this.modalVisible = false;
          this.fetchPermissions();
        });
      },
      handleEditPermission() {
        this.$api.permission.updatePermission(this.permission.id, this.permission).then(() => {
          this.$message.success('更新权限成功');
          this.modalVisible = false;
          this.fetchPermissions();
        });
      },
      permTypeFormat(permType) {
        let result;
        switch (permType) {
          case 0:
            result = `<span class="text-warning">操作</span>`;
            break;
          case 1:
            result = `<span class="text-info">菜单</span>`;
            break;
          case 2:
            result = `<span class="text-success">按钮</span>`;
            break;
          default:
            result = `<span class="text-mute">未知</span>`;
        }
        return result;
      },
      handlePagerChange(page, pageSize) {
        this.pager.current = page;
        this.pager.pageSize = pageSize;
        this.fetchPermissions();
      }
    }
  }
</script>

<style scoped>

</style>