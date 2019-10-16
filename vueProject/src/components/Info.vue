<template>
  
  <div style="margin-top:3em">
    <el-row>
      <el-col :span="15" :offset="4">
        <el-card>
          <div slot="header" class="clearfix" style="text-align:left">
            <span>概览</span>
          </div>
          <el-col :span="23" :offset="1">
            <el-form label-width="80px" :model="user">
              <el-form-item label="帐号">
                <el-col :span="20">
                  <el-input v-model="user.username" disabled></el-input>
                </el-col>
              </el-form-item>
              <el-form-item label="昵称">
                <el-col :span="20">
                  <el-input v-model="user.nickname" :disabled="nicknameDisable"></el-input>
                </el-col>
                <el-col :span="2" :offset="1">
                  <el-button v-if="nicknameDisable" type="text" @click="editNickName">修改</el-button>
                  <el-button v-else type="text" @click="saveNickName">保存</el-button>
                </el-col>
              </el-form-item>
              <el-form-item label="描述">
                <el-col :span="20">
                  <el-input type="textarea" v-model="user.description" :disabled="descriptionDisable"></el-input>
                </el-col>
                <el-col :span="2" :offset="1">
                  <el-button v-if="descriptionDisable" type="text" @click="editDescription">修改</el-button>
                  <el-button v-else type="text" @click="saveDescription">保存</el-button>
                </el-col>
              </el-form-item>
            </el-form>
          </el-col>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {

      state:"",
      nicknameDisable: true,
      descriptionDisable:true,
      user: {
        username: "",
        nickname: "",
        description: "",
        avatar: "",
      }
    };
  },
  methods: {
    editNickName() {
      this.nicknameDisable = !this.nicknameDisable;
    },
    saveNickName() {
      this.axios.post(this.BaseUrl + "/api/v1/user/nickname",{"nickname":this.user.nickname}).then(res=>{
        var data = res.data
        if(data.status == 66666){
          this.$store.commit("setUserNickname",this.user.nickname)
          this.editNickName()
        }else{
          this.$message.error(data.data)
        }
      })
    },
    editDescription() {
      this.descriptionDisable = !this.descriptionDisable;
    },
    saveDescription() {
      this.axios.post(this.BaseUrl + "/api/v1/user/description",{"description":this.user.description}).then(res=>{
        var data = res.data
        if(data.status == 77777){
          this.$store.commit("setUserDescription",this.user.description)
          this.editDescription()
        }else{
          this.$message.error(data.data)
        }
      })
    }
  },
  mounted() {
   
    var state = this.$store.state
    this.user = state.user
  
  },
  computed: {
    init(){
      return this.$store.state.user;
    },
  },
};
</script>

<style>
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}
</style>