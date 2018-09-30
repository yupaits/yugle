<template>
  <div>
    <h3>控制台</h3>
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
    height: 480px;
  }
</style>