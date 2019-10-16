<template>
  <el-row style="margin-top:5em">
    <el-col :span="8" :offset="8">
      <el-card :body-style="{ padding: '0px' }">
        <div slot="header" style="text-align:left">
          <span>卡片名称</span>
        </div>
        <el-form :model="form" :rules="rules" ref="signInForm" label-width="100px" style="padding-right:3em;margin-top:20px">
          <el-form-item label="帐号:" prop="username">
            <el-input v-model="form.username"></el-input>
          </el-form-item>
          <el-form-item label="昵称:" prop="nickname">
            <el-input v-model="form.nickname"></el-input>
          </el-form-item>
          <el-form-item label="密码:" prop="password">
            <el-input v-model="form.password"></el-input>
          </el-form-item>
          <el-form-item label="确认密码:" prop="password_confirm">
            <el-input v-model="form.password_confirm"></el-input>
          </el-form-item>
          <el-form-item label>
            <el-button type="primary" @click="onSubmit">注册</el-button>
            <router-link to="/">
              <el-button type="primary">回到主页</el-button>
            </router-link>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>
<script>
export default {
  data() {
    var validatePass2 = (rule, value, callback) => {
      if (value != this.form.password) {
        callback(new Error("两次输入密码不一致"));
      } else {
        callback();
      }
    };
    return {
      form: {
        username: "",
        password: "",
        nickname: "",
        password_confirm: ""
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          { min: 3, max: 16, message: "长度在 3 到 16 个字符", trigger: "blur" }
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 3, max: 16, message: "长度在 3 到 16 个字符", trigger: "blur" }
        ],
        nickname: [
          { required: true, message: "请输入昵称", trigger: "blur" },
          { min: 3, max: 16, message: "长度在 3 到 16 个字符", trigger: "blur" }
        ],
        password_confirm: [
          { required: true, message: "请确认密码", trigger: "blur" },
          {
            min: 3,
            max: 16,
            message: "长度在 3 到 16 个字符",
            trigger: "blur"
          },
          { validator: validatePass2, trigger: "blur" }
        ]
      }
    };
  },

  methods: {
    onSubmit() {
      this.$refs["signInForm"].validate(vaild => {
        if (vaild) {
          this.axios
            .post(this.BaseUrl + "/api/v1/signIn", this.form)
            .then(res => {
              console.log(res);
            })
            .catch(e => {
              console.log(e);
            });
        } else {
          alert("!");
        }
      });
    }
  }
};
</script>