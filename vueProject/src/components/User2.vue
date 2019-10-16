<template>
  <div style="margin-top:3em">
    <el-row>
      <el-col :span="10" :offset="7">
        <el-card class="box-card">
          <div slot="header" class="clearfix" style="text-align:left">
            <span>个人信息</span>
            <el-button style="float: right; padding: 3px 0" type="text">保存</el-button>
          </div>
          <el-form
            :model="form"
            ref="signInForm"
            label-width="100px"
            style="padding-right:3em;margin-top:20px"
          >
            <el-form-item label="头像:" style="text-align:left">
              <el-avatar :size="60" :src="form.avatar"></el-avatar><el-link :underline="false" type="primary" href="/face" style="margin-left:40px" >11</el-link>
            </el-form-item>
            <el-form-item label="帐号:" prop="username">
              <el-input v-model="form.username" disabled></el-input>
            </el-form-item>
            <el-form-item label="昵称:" prop="nickname">
              <el-input v-model="form.nickname"></el-input>
            </el-form-item>
            <el-form-item label="描述:" prop="description">
              <el-input v-model="form.description"></el-input>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        id: "",
        username: "",
        nickname: "",
        avatar: "",
        description: ""
      },
      name: "user"
    };
  },
  mounted() {
    var user = this.$store.state.user;
    console.log(user);
    console.log(this.$store.state)
    this.form.username = user.username;
    this.form.nickname = user.nickname;
    this.form.avatar = user.avatar;
    this.form.description = user.description;
  }
};
</script>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>