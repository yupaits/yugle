<template>
  <a-layout>
    <a-layout-header>
      <span class="logo mr-3">
        <img src="/favicon.ico" alt="Logo" width="42" height="42" class="logo-img">
        <span class="logo-title ml-1" @click="$router.push('/dashboard')">鱼钩资讯</span>
      </span>
      <span class="header-user">
        <a-popover trigger="click" placement="bottomRight">
          <span class="avatar">
            <a-badge :dot="true">
              <a-avatar src="https://blog.yupaits.com/avatar.jpg" shape="circle" icon="user"/>
            </a-badge>
          </span>
          <template slot="content">
            <div v-for="item in $consts.userMenu" :key="item.label">
              <router-link :to="item.to" class="user-menu-item"><a-icon :type="item.icon"/> {{item.label}}</router-link>
            </div>
            <a href="javascript:;" class="user-menu-item" @click="logout"><a-icon type="logout"/> 注销登录</a>
          </template>
        </a-popover>
      </span>
      <a-menu
          theme="dark"
          mode="horizontal"
          :selectedKeys="[$store.getters.key]"
          :style="{ lineHeight: '64px' }">
        <a-sub-menu v-for="nav in $consts.navs" :key="nav.key">
          <span slot="title">
            <i class="iconfont" :class="nav.icon"/> {{nav.label}}
          </span>
          <a-menu-item v-for="item in nav.items" :key="item.key" @click.native="() => {$router.push(item.to)}">
            {{item.label}}
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-header>
    <a-layout-content class="layout-content-container">
      <a-layout class="layout-content">
        <router-view/>
      </a-layout>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      yugle 鱼钩资讯 ©2018 Created by <a href="https://github.com/YupaiTS" target="_blank">YupaiTS</a>
    </a-layout-footer>
  </a-layout>
</template>

<script>
  import 'vue-cookies'
  export default {
    name: "Index",
    methods: {
      logout() {
        window.$cookies.remove(this.$consts.token.cookieName);
        window.location.href = '/login';
      }
    }
  }
</script>

<style scoped>
  .ant-layout-header {
    height: 65px;
    line-height: 60px;
  }

  .logo {
    float: left;
  }

  .header-user {
    float: right;
  }

  .logo-img {
    border-radius: 8px;
  }

  .logo-title {
    font-size: 22px;
    color: #d9d9d9;
  }

  .logo-title:hover {
    cursor: pointer;
    color: #fafafa;
  }

  .avatar:hover {
    cursor: pointer;
  }

  .user-menu-item {
    line-height: 36px;
    color: #262626;
  }

  .user-menu-item:hover {
    color: #1890ff;
  }

  .layout-content-container {
    padding: 36px 36px 0 36px;
  }

  .layout-content {
    background: #fff;
    min-height: calc(100vh - 170px);
    padding: 24px;
  }
</style>