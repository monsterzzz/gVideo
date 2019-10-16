 

<template>
  <div style="margin-top: 5em">
    <el-row>
      <el-col :span="8" :offset="8">
        <el-card :body-style="{ padding: '0px' }">
          <el-image style="height:160px; overflow:hidden" :src="imgSrc" fit="fill"></el-image>
          <div style="margin-top: 20px;padding-right:3em;padding-left:1em">
            <el-form ref="form" :model="form" label-width="80px">
              <el-form-item label="用户名:">
                <el-input v-model="form.username"></el-input>
              </el-form-item>
              <el-form-item label="密码:">
                <el-input v-model="form.password"></el-input>
              </el-form-item>
              <el-form-item label-width="0px">
                <el-button type="primary" @click="onSubmit">登录</el-button>
                <router-link to="/signIn" style="margin-left:20px">
                  <el-button>注册</el-button>
                </router-link>
              </el-form-item>
            </el-form>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  data() {
    return {
      fit: "fill",
      form: {
        username: "",
        password: ""
      },
      ipt: "",
      msg: "Welcome to Your Vue.js App",
      imgSrc: "/static/9.png"
    };
  },

  // created() {
  //   console.log(this.$store.state);
  // },
  methods: {
    onSubmit() {
      
      this.axios
        //.post("/api/api/v1/login", this.form)
        .post(this.BaseUrl + "/api/v1/login", this.form)
        .then(response => {
  
          if(response.data.status == 99999){
            this.$message.error("error form")
          }else{
            var data = response.data.data
            data.avatar = this.BaseUrl + "" + data.avatar
            var a = JSON.stringify(data)
            this.$store.commit("setUser", response.data.data);
            this.$router.push({ path: "/main" });
          }
        })
        .catch(e => {
          console.log(e);
        });
    },
    
    onSignIn() {
      this.axios
        .post("/api/api/v1/signIn", this.form)
        .then(response => {
          console.log(response.data);
        })
        .catch(e => {
          console.log(e);
        });
      // console.log(this.$store)
      // this.$store.commit('setToken',"helloWorld")
      // console.log(this.$store.state.token)
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #ebebeb;
}
.login {
  margin-top: 30px;
  border-radius: 30px;
  background-color: rgb(247, 244, 244);

  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  -webkit-mask-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAA5JREFUeNpiYGBgAAgwAAAEAAGbA+oJAAAAAElFTkSuQmCC);
}

.login-row {
}
</style>
