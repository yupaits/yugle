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
             :scroll="$consts.tableScroll.page"></a-table>
    <a-pagination size="small"
                  class="pager"
                  v-model="pager.current"
                  :total="pager.total"
                  :pageSize.sync="pager.pageSize"
                  :pageSizeOptions="pager.pageSizeOptions"
                  :showQuickJumper="pager.showQuickJumper"
                  :showSizeChanger="pager.showSizeChanger"
                  @change="handlePagerChange"></a-pagination>
  </div>
</template>

<script>
  const columns = [
    {title: '用户名', width: '', dataIndex: 'Username'},
    {title: '性别', width: '', scopedSlots: {customRender: ''}},
    {title: '手机号码', width: '', dataIndex: 'Phone'},
    {title: 'Email', width: '', dataIndex: 'Email'},
    {title: '出生日期', width: '', scopedSlots: {customRender: ''}},
    {title: '状态', width: '', scopedSlots: {customRender: ''}},
    {title: '操作', width: '', scopedSlots: {customRender: ''}},
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