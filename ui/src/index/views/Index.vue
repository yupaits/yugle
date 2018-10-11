<template>
  <a-layout style="min-height: 100vh">
    <a-layout-sider collapsible v-model="collapsed">
      <div class="logo">
        <span class="pull-left">
          <img src="/favicon.ico" alt="Logo" :width="collapsed ? 36 : 32" :height="collapsed ? 36 : 32" class="logo-img">
        </span>
        <span class="logo-title" v-if="!collapsed">鱼钩资讯</span>
      </div>
      <a-menu theme="dark" mode="inline" :selectedKeys="[$store.getters.key]">
        <a-menu-item key="dashboard" @click.native="$router.push('/dashboard')">
          <a-icon type="dashboard"/><span>控制台</span>
        </a-menu-item>
        <a-sub-menu v-for="nav in $consts.navs" :key="nav.key">
          <span slot="title">
            <a-icon :type="nav.icon"/><span>{{nav.label}}</span>
          </span>
          <a-menu-item v-for="item in nav.items" :key="item.key" @click.native="$router.push(item.to)">
            <a-icon :type="nav.icon"/>{{item.label}}
          </a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="header">
        <span class="pull-right">
          <a-popover v-model="userMenuVisible" trigger="click" placement="bottomRight">
            <span class="avatar">
              <a-badge :dot="true">
                <a-avatar src="https://blog.yupaits.com/avatar.jpg" shape="circle" icon="user"/>
              </a-badge>
            </span>
            <template slot="content">
              <div v-for="item in $consts.userMenu" :key="item.label" class="user-menu-item" @click="goPage(item.to)">
                <a-icon :type="item.icon"/> {{item.label}}
              </div>
              <div class="user-menu-item" @click="logout"><a-icon type="logout"/> 注销登录</div>
            </template>
          </a-popover>
        </span>
      </a-layout-header>
      <a-layout-content>
        <router-view/>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        yugle 鱼钩资讯 ©2018 Created by <a href="https://github.com/YupaiTS" target="_blank">YupaiTS</a>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
  import 'vue-cookies'
  export default {
    name: "Index",
    data() {
      return {
        collapsed: false,
        userMenuVisible: false
      }
    },
    methods: {
      goPage(path) {
        this.userMenuVisible = false;
        this.$router.push(path);
      },
      logout() {
        window.$cookies.remove(this.$consts.token.cookieName);
        window.location.href = '/login';
      }
    }
  }
</script>

<style scoped>
  .ant-layout-header {
    line-height: 58px;
  }

  .header {
    padding: 0 16px;
    background: #ffffff;
    border-bottom: 1px solid #d9d9d9;
  }

  .logo {
    height: 32px;
    margin: 16px;
  }

  .logo-img {
    border-radius: 8px;
  }

  .logo-title {
    color: #d9d9d9;
    font-size: 22px;
    margin-left: 8px;
  }

  .avatar:hover {
    cursor: pointer;
  }

  .user-menu-item {
    line-height: 36px;
    color: #262626;
  }

  .user-menu-item:hover {
    cursor: pointer;
    color: #1890ff;
  }
</style>