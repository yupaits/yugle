<template>
  <div id="app">
    <div class="container">
      <div class="row">
        <div class="col-4"></div>
        <div class="col-4" style="margin-top: 6rem;">
          <div class="card" style="width: 24rem;">
            <div class="card-body">
              <h5 class="card-title">
                <img src="/favicon.ico" alt="logo"> 登录
              </h5>
              <div class="alert alert-danger" v-if="showAlert">
                <span>{{alertInfo}}</span>
              </div>
              <form class="form-group">
                <div class="mb-3">
                  <span class="label label-important">用户名：</span>
                  <input type="text" v-model="loginForm.username" class="form-control form-control-sm" placeholder="请输入用户名"/>
                </div>
                <div class="mb-3">
                  <span class="label label-important">密码：</span>
                  <input type="password" v-model="loginForm.password" class="form-control form-control-sm" placeholder="请输入密码"/>
                </div>
                <button class="btn btn-primary btn-sm" @click="login">登 录</button>
              </form>
            </div>
          </div>
        </div>
        <div class="col-4"></div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        showAlert: false,
        alertInfo: '',
        loginForm: {
          username: '',
          password: ''
        }
      }
    },
    methods: {
      login() {
        this.auth.post('/login', this.loginForm).then(res => {
          alert(JSON.stringify(res));
        }).catch(error => {
          alert(JSON.stringify(error));
        });
      }
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }
</style>
