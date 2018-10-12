<template>
  <div>
    <a-row>
      <a-col :span="12" :push="6">
        <a-form>
          <a-form-item label="当前密码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol" required
                       :validateStatus="validate.oldPassword.status" :help="validate.oldPassword.help">
            <a-input type="password" v-model="passwordModify.oldPassword" placeholder="请填写当前密码"></a-input>
          </a-form-item>
          <a-form-item label="新的密码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol" required
                       :validateStatus="validate.newPassword.status" :help="validate.newPassword.help">
            <a-input type="password" v-model="passwordModify.newPassword" placeholder="请填写新密码"></a-input>
          </a-form-item>
          <a-form-item label="确认密码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol" required
                       :validateStatus="validate.confirmPassword.status" :help="validate.confirmPassword.help">
            <a-input type="password" v-model="passwordModify.confirmPassword" placeholder="重复输入确认新密码"></a-input>
          </a-form-item>
          <a-form-item :wrapperCol="$consts.form.submitCol">
            <a-button type="primary" :loading="btnLoading" @click="handleModify">确认修改</a-button>
          </a-form-item>
        </a-form>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    name: "ModifyPassword",
    data() {
      return {
        passwordModify: {
          oldPassword: '',
          newPassword: '',
          confirmPassword: ''
        },
        validate: {
          oldPassword: {
            status: '',
            help: ''
          },
          newPassword: {
            status: '',
            help: ''
          },
          confirmPassword: {
            status: '',
            help: ''
          }
        },
        btnLoading: false
      }
    },
    methods: {
      handleModify() {
        if (this.validatePasswordModify()) {
          this.btnLoading = true;
          this.$api.user.modifyPassword(this.passwordModify).then(() => {
            this.$message.success('修改密码成功');
            this.passwordModify = {
              oldPassword: '',
              newPassword: '',
              confirmPassword: ''
            };
            this.btnLoading = false;
          }).catch(() => {
            this.btnLoading = false;
          });
        }
      },
      validatePasswordModify() {
        let valid = true;
        if (this.passwordModify.oldPassword) {
          this.validate.oldPassword.status = '';
          this.validate.oldPassword.help = '';
        } else {
          this.validate.oldPassword.status = 'error';
          this.validate.oldPassword.help = '当前密码不能为空';
          valid = false;
        }
        if (this.passwordModify.newPassword) {
          this.validate.newPassword.status = '';
          this.validate.newPassword.help = '';
        } else {
          this.validate.newPassword.status = 'error';
          this.validate.newPassword.help = '新的密码不能为空';
          valid = false;
        }
        if (this.passwordModify.confirmPassword) {
          if (this.passwordModify.confirmPassword !== this.passwordModify.newPassword) {
            this.validate.confirmPassword.status = 'error';
            this.validate.confirmPassword.help = '两次输入的密码不一致';
            valid = false;
          } else {
            this.validate.confirmPassword.status = '';
            this.validate.confirmPassword.help = '';
          }
        } else {
          this.validate.confirmPassword.status = 'error';
          this.validate.confirmPassword.help = '确认密码不能为空';
          valid = false;
        }
        return valid;
      }
    }
  }
</script>

<style scoped>

</style>