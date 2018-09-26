<template>
  <div>
    <a-row class="opt-container">
      <a-button icon="plus" class="btn-success">添加用户</a-button>
      <div class="pull-right">
        <span>关键字：
          <a-input v-model="userQuery.keyword" placeholder="请输入用户名关键字" class="search-input mr-2"></a-input>
        </span>
        <span>状态：
          <a-select v-model="userQuery.enabled" class="mr-2" @change="handleSearch">
            <a-select-option :value="true">激活</a-select-option>
            <a-select-option :value="false">禁用</a-select-option>
          </a-select>
        </span>
        <a-button icon="search" @click="handleSearch">搜索</a-button>
      </div>
    </a-row>
    <a-table size="small"
             :columns="columns"
             :rowKey="record => record.ID"
             :dataSource="users"
             :pagination="false"
             :loading="loading"
             :scroll="$consts.tableScroll.page">
      <template slot="user" slot-scope="record">
        <a-avatar shape="square" :src="record.Avatar" icon="user"/>
        <span class="ml-1">{{record.Username}}</span>
      </template>
      <template slot="gender" slot-scope="record">
        <a-tag v-if="record.Gender === 0">
          <a-icon type="woman"/> 女
        </a-tag>
        <a-tag v-else-if="record.Gender === 1">
          <a-icon type="man"/> 男
        </a-tag>
        <a-tag v-else-if="record.Gender === 2">
          <a-icon type="user"/> 保密
        </a-tag>
      </template>
      <template slot="birthday" slot-scope="record">
        {{record | birthdayFormat}}
      </template>
      <template slot="status" slot-scope="record">
        <a-tag color="#389e0d" v-if="record.Enabled">激活</a-tag>
        <a-tag color="#f5222d" v-else>禁用</a-tag>
      </template>
      <template slot="operation" slot-scope="record">
        <a-popconfirm title="确定要禁用该用户吗？" placement="topRight"
                      @confirm="disableUser(record.UserId)" okText="确定" cancelText="取消" v-if="record.Enabled">
          <a-button type="danger" size="small">禁用</a-button>
        </a-popconfirm>
        <a-popconfirm title="确定要解禁该用户吗？" placement="topRight"
                      @confirm="enableUser(record.UserId)" okText="确定" cancelText="取消" v-else>
          <a-button size="small" class="btn-success">解禁</a-button>
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

    <a-modal></a-modal>
  </div>
</template>

<script>
  const columns = [
    {title: '用户名', width: '15%', scopedSlots: {customRender: 'user'}},
    {title: '性别', width: '10%', scopedSlots: {customRender: 'gender'}},
    {title: '手机号码', width: '15%', dataIndex: 'Phone'},
    {title: 'Email', width: '15%', dataIndex: 'Email'},
    {title: '出生日期', width: '15%', scopedSlots: {customRender: 'birthday'}},
    {title: '状态', width: '15%', scopedSlots: {customRender: 'status'}},
    {title: '操作', width: '15%', scopedSlots: {customRender: 'operation'}},
  ];
  export default {
    name: "UserManage",
    data() {
      return {
        columns: columns,
        users: [],
        pager: {
          showQuickJumper: true,
          showSizeChanger: true,
          current: 1,
          total: 1,
          pageSize: parseInt(this.$consts.pageSizes[0]),
          pageSizeOptions: this.$consts.pageSizes
        },
        loading: false,
        userQuery: {
          keyword: '',
          enabled: true
        }
      }
    },
    mounted() {
      this.fetchUsers();
    },
    filters: {
      birthdayFormat(birth) {
        return birth.BirthYear + '-' + birth.BirthMonth + '-' + birth.BirthDay;
      }
    },
    methods: {
      fetchUsers() {
        this.loading = true;
        this.$api.user.getUserPage(this.pager.current, this.pager.pageSize, this.userQuery).then(res => {
          this.users = res.data.Content;
          this.pager.current = res.data.Page;
          this.pager.total = res.data.Total;
          this.loading = false;
        }).catch(() => {
          this.loading = false;
        });
      },
      handleSearch() {
        this.fetchUsers();
      },
      disableUser(userId) {
        this.$api.user.changeUserStatus(userId, false).then(() => {
          this.$message.success('禁用用户成功');
          this.fetchUsers();
        });
      },
      enableUser(userId) {
        this.$api.user.changeUserStatus(userId, true).then(() => {
          this.$message.success('禁用用户成功');
          this.fetchUsers();
        });
      },
      handlePagerChange(page, pageSize) {
        this.pager.current = page;
        this.pager.pageSize = pageSize;
        this.fetchUsers();
      }
    }
  }
</script>

<style scoped>

</style>