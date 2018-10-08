<template>
  <div>
    <h2>个人中心</h2>
    <a-row>
      <a-col :span="12" :push="6">
        <a-form>
          <a-form-item label="头像" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-avatar :src="avatar" :size="80" icon="user"></a-avatar>
            <div class="mt-1">
              <a-upload action=""
                        accept="image/*"
                        :showUploadList="false"
                        @change="handleAvatarChange">
                <a-button>浏览文件...</a-button>
              </a-upload>
              <a-popconfirm title="确定要删除头像吗？" placement="topLeft"
                            @confirm="removeAvatar" okText="确定" cancelText="取消">
                <a-button class="ml-3">删除头像</a-button>
              </a-popconfirm>
            </div>
          </a-form-item>
          <a-form-item label="用户名" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-input :value="user.username" disabled class="form-input"></a-input>
          </a-form-item>
          <a-form-item label="手机号码" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-input v-model="user.phone" placeholder="请填写手机号码" class="form-input"></a-input>
          </a-form-item>
          <a-form-item label="电子邮箱" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-input v-model="user.email" placeholder="请填写电子邮箱" class="form-input"></a-input>
          </a-form-item>
          <a-form-item label="性别" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-radio-group v-model="user.gender">
              <a-radio :value="0">女</a-radio>
              <a-radio :value="1">男</a-radio>
              <a-radio :value="2">保密</a-radio>
            </a-radio-group>
          </a-form-item>
          <a-form-item label="出生日期" :labelCol="$consts.form.labelCol" :wrapperCol="$consts.form.wrapperCol">
            <a-date-picker format="YYYY年MM月DD日" :defaultValue="moment('2000-01-01')" :showToday="false"
                           @change="handleBirthChange"></a-date-picker>
          </a-form-item>
          <a-form-item :wrapperCol="$consts.form.submitCol">
            <a-button class="btn-success" :loading="btnLoading" @click="save">更新个人资料</a-button>
          </a-form-item>
        </a-form>
      </a-col>
    </a-row>
  </div>
</template>

<script>
  import moment from 'moment'
  export default {
    name: "Profile",
    data() {
      return {
        avatar: '',
        user: {},
        btnLoading: false,
      }
    },
    mounted() {
      this.fetchUser();
    },
    methods: {
      moment,
      fetchUser() {
        this.$api.user.getCurrentUser().then(res => {
          this.user = {
            username: res.data.Username,
            userId: res.data.UserId,
            phone: res.data.Phone,
            email: res.data.Email,
            gender: res.data.Gender,
            avatar: res.data.Avatar,
            birthYear: res.data.BirthYear,
            birthMonth: res.data.BirthMonth,
            birthDay: res.data.BirthDay,
          };
        });
      },
      handleAvatarChange(info) {

      },
      removeAvatar() {
        this.avatar = '';
        this.user.avatar = '';
      },
      handleBirthChange(date) {
        this.user.birthYear = date.year();
        this.user.birthMonth = date.month() + 1;
        this.user.birthDay = date.day();
      },
      save() {
        this.btnLoading = true;
        this.$api.user.updateUser(this.user.userId, this.user).then(() => {
          this.$message.success('更新个人资料成功');
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