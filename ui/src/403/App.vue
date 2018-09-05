<template>
  <div id="app">
    <img src="https://gw.alipayobjects.com/zos/rmsportal/wZcnGqRDyhPOEYFcZDnb.svg" alt="">
    <h1 class="title">403</h1>
    <h3 class="sub-title text-mute">抱歉，你无权访问该页面</h3>
    <a-button type="primary" @click="gotoIndex">返回主页</a-button>
  </div>
</template>

<script>
  export default {
    methods: {
      gotoIndex() {
        window.location.href = '/';
      }
    }
  }
</script>

<style>
  #app {
    font-family: Chinese Quote, -apple-system, BlinkMacSystemFont, Segoe UI, PingFang SC, Hiragino Sans GB, Microsoft YaHei,
    Helvetica Neue, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
    background: #f0f2f5;
    padding-top: 100px;
    min-height: calc(100vh);
    text-align: center;
  }

  .title {
    color: #434e59;
    font-size: 72px;
    font-weight: 600;
    line-height: 72px;
    margin: 24px 0;
  }

  .sub-title {
    font-size: 20px;
    line-height: 28px;
    margin: 16px 0;
  }
</style>
