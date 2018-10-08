<template>
  <div>
    <h2>系统设置</h2>
    <a-row :gutter="16">
      <a-col :span="3">
        <a-menu theme="light" mode="inline" :selectedKeys="$store.getters.settingsKeys" class="settings-menu">
          <a-menu-item v-for="item in $consts.settingsMenu" :key="item.key" @click.native="goPage(item)">{{item.label}}</a-menu-item>
        </a-menu>
      </a-col>
      <a-col :span="21">
        <router-view/>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    name: "Settings",
    methods: {
      goPage(item) {
        this.$router.push('/settings/' + item.key);
      }
    }
  }
</script>

<style scoped>
  .settings-menu {
    min-height: calc(100vh - 249px);
  }
</style>