<template>
  <div>
    <h2>个人中心</h2>
    <a-row>
      <a-col :span="12" :push="6">
        <a-form>
          <a-form-item label="头像" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-avatar :size="64" icon="user"></a-avatar>
          </a-form-item>
          <a-form-item label="用户名" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-input disabled class="form-input"></a-input>
          </a-form-item>
          <a-form-item label="手机号码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-input placeholder="请填写手机号码" class="form-input"></a-input>
          </a-form-item>
          <a-form-item label="电子邮箱" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-input placeholder="请填写电子邮箱" class="form-input"></a-input>
          </a-form-item>
          <a-form-item :wrapperCol="$consts.form.submitCol">
            <a-button class="btn-success">更新个人资料</a-button>
          </a-form-item>
        </a-form>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  export default {
    name: "Profile"
  }
</script>

<style scoped>

</style>