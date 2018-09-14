<template>
  <a-card>
    <h3 class="mb-1"><a-icon type="clock-circle-o"/> 定时任务</h3>
    <a-table size="small"
             :columns="columns"
             :rowKey="record => record.ID"
             :dataSource="tasks"
             :pagination="false"
             :loading="loading"
             :scroll="$consts.tableScroll.card">
      <template slot="state" slot-scope="record">
        <a-tag color="#ffa940" v-if="record.State === 0">空闲</a-tag>
        <a-tag color="#389e0d" v-else-if="record.State === 1">执行中</a-tag>
        <a-tag color="#d9d9d9" v-else>未知</a-tag>
      </template>
      <template slot="next" slot-scope="record">
        {{record.Next | goTimeFormat}}
      </template>
      <template slot="prev" slot-scope="record">
        {{record.Prev | goTimeFormat}}
      </template>
      <template slot="last-result" slot-scope="record">
        <div v-if="$consts.golangDateTimeFormat(record.Prev)">
          <a-tag color="#87d068" v-if="record.LastSuccess"><a-icon type="check"/> 成功</a-tag>
          <a-tag color="#f50" v-else><a-icon type="close"/> 失败</a-tag>
        </div>
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
  </a-card>
</template>

<script>
  const columns = [
    {title: '状态', width: '15%', scopedSlots: {customRender: 'state'}},
    {title: '任务名称', width: '30%', dataIndex: 'TaskName'},
    {title: '下次执行时间', width: '20%', scopedSlots: {customRender: 'next'}},
    {title: '上次执行时间', width: '20%', scopedSlots: {customRender: 'prev'}},
    {title: '上次执行结果', width: '15%', scopedSlots: {customRender: 'last-result'}},
  ];
  export default {
    name: "TaskMonitor",
    data() {
      return {
        columns: columns,
        tasks: [],
        pager: {
          showQuickJumper: true,
          showSizeChanger: true,
          current: 1,
          total: 1,
          pageSize: parseInt(this.$consts.pageSizes[0]),
          pageSizeOptions: this.$consts.pageSizes
        },
        loading: false,
      }
    },
    mounted() {
      this.fetchTasks();
    },
    methods: {
      fetchTasks() {
        this.loading = true;
        this.$api.task.getTasks(this.pager.current, this.pager.pageSize).then(res => {
          this.tasks = res.data.Content;
          this.pager.current = res.data.Page;
          this.pager.total = res.data.Total;
          this.loading = false;
        }).catch(() => {
          this.loading = false;
        });
      },
      handlePagerChange(page, pageSize) {
        this.pager.current = page;
        this.pager.pageSize = pageSize;
        this.fetchTasks();
      }
    }
  }
</script>

<style scoped>

</style>