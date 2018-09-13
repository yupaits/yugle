<template>
  <div>
    <h3>用户管理</h3>
  </div>
</template>

<script>
  const columns = [];
  export default {
    name: "UserManage",
    data() {
      return {
        columns: columns,
        users: []
      }
    },
    mounted() {
      this.fetchUsers();
    },
    methods: {
      fetchUsers() {
        this.$api.user.getUserPage().then().catch(() => {

        });
      }
    }
  }
</script>

<style scoped>

</style>