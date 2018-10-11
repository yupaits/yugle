<template>
  <div>
    <a-breadcrumb class="breadcrumb-container">
      <a-breadcrumb-item>
        <a-icon type="home"/>
        <span>主页</span>
      </a-breadcrumb-item>
      <a-breadcrumb-item v-for="item in items" :key="item">{{item}}</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
</template>

<script>
  export default {
    name: "Breadcrumb",
    props: {
      items: {
        type: Array,
        required: true
      }
    }
  }
</script>

<style scoped>
  .breadcrumb-container {
    padding: 16px;
  }
</style>