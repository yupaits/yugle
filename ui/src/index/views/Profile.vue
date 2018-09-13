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
            <a-input :value="username" disabled class="form-input"></a-input>
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
            <a-button class="btn-success" @click="save">更新个人资料</a-button>
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
        username: 'admin',
        avatar: 'https://blog.yupaits.com/avatar.jpg',
        user: {
          userId: '',
          phone: '',
          email: '',
          gender: 2,
          avatar: '',
          birthYear: undefined,
          birthMonth: undefined,
          birthDay: undefined,
        }
      }
    },
    methods: {
      moment,
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
        alert(JSON.stringify(this.user));
      }
    }
  }
</script>

<style scoped>

</style>