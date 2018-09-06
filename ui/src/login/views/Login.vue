<template>
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
                <input type="text" v-model="loginForm.username" class="form-control form-control-sm"
                       placeholder="请输入用户名" @keyup.enter="login"/>
              </div>
              <div class="mb-3">
                <span class="label label-important">密码：</span>
                <input type="password" v-model="loginForm.password" class="form-control form-control-sm"
                       placeholder="请输入密码" @keyup.enter="login"/>
              </div>
              <button type="button" class="btn btn-primary btn-sm" @click="login">登 录</button>
            </form>
          </div>
        </div>
      </div>
      <div class="col-4"></div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Login",
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
        if (this.validateLoginForm()) {
          this.auth.post('/login', this.loginForm).then(res => {
            this.$cookies.set(this.consts.token.cookieName, res.token, this.consts.token.maxRefreshTime);
            this.$cookies.set(this.consts.token.refreshFlag, 'flag', this.consts.token.expiredTime);
            window.location.href = '/';
          });
        }
      },
      validateLoginForm() {
        let valid = true;
        if (this.loginForm.username === '' || this.loginForm.password === '') {
          this.showAlert = true;
          this.alertInfo = '用户名和密码都不能为空';
          valid = false;
        } else {
          this.showAlert = false;
        }
        return valid;
      }
    }

  }
</script>

<style scoped>

</style>