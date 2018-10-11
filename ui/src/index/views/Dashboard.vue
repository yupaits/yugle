<template>
  <div>
    <breadcrumb :items="['控制台']"/>
    <div class="layout-content">
      <a-row :gutter="16">
        <a-col :span="8">

        </a-col>
        <a-col :span="16">
          <task-monitor class="monitor-card"></task-monitor>
        </a-col>
      </a-row>
    </div>
  </div>
</template>

<script>
  import TaskMonitor from "./monitor/TaskMonitor";
  import Breadcrumb from "../../components/Breadcrumb";
  export default {
    name: "Dashboard",
    components: {TaskMonitor, Breadcrumb}
  }
</script>

<style scoped>
  .monitor-card {
    height: 480px;
  }
</style>