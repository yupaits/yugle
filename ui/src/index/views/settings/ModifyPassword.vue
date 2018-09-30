<template>
  <div>
    <a-row>
      <a-col :span="12" :push="6">
        <a-form>
          <a-form-item label="当前密码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol" required>
            <a-input type="password" v-model="passwordModify.oldPassword" placeholder="请填写当前密码"></a-input>
          </a-form-item>
          <a-form-item label="新的密码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol" required>
            <a-input type="password" v-model="passwordModify.newPassword" placeholder="请填写新密码"></a-input>
          </a-form-item>
          <a-form-item label="确认密码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol" required>
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
        btnLoading: false
      }
    },
    methods: {
      handleModify() {
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
    }
  }
</script>

<style scoped>

</style>