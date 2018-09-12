<template>
  <div>
    <h2>监控面板</h2>
    <a-row :gutter="16">
      <a-col :span="8">

      </a-col>
      <a-col :span="16">
        <task-monitor class="monitor-card"></task-monitor>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  import TaskMonitor from "./monitor/TaskMonitor";
  export default {
    name: "Dashboard",
    components: {TaskMonitor}
  }
</script>

<style scoped>
  .monitor-card {
    height: 460px;
  }
</style>