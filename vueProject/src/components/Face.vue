<template>
  <div>
    <el-row style="margin-top:5em">
      <el-col :span="3" :offset="8">
        <el-upload ref="avatar"
          style="text-align:left"
          class="avatar-uploader"
          action="/api/api/v1/image/avatar/upload"
          :on-change="onChange"
          :auto-upload="false"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
        >
          <img v-if="imageUrl" :src="imageUrl" class="avatar" />
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-col>
      <el-col :span="1">
        <div style="text-align:left;background-color:black;width:1px;height:178px;"></div>
      </el-col>
      <el-col :span="2" style="text-align:left">
         <el-avatar :size="70" :src="imageUrl" style="margin-top:3em"></el-avatar>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="6" :offset="8" style="margin-top:40px">
        <el-button @click="submit">确定</el-button>
      </el-col>
    </el-row>
  </div>
</template>

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

<script>
export default {
  data() {
    return {
      imageUrl: ""
    };
  },
  methods: {
    handleAvatarSuccess(res, file) {
      console.log(res)
      this.$message.info("上传成功!")
    },
    submit(){
      var f = new FormData()
      console.log(this.$refs.avatar.uploadFiles)
      f.append("file",this.$refs.avatar.uploadFiles[0].raw)
      this.axios.post("http://localhost:8000/api/v1/image/avatar/upload",f).then(res=>{
        var data = res.data
        if(data.status == 50000){
          console.log(data.data)
          this.$store.commit("setUserAvatar","http://localhost:8000" + data.data.path)
        }
        console.log(this.$store.state)
      })
    },
    onChange(file) {
      
      const isJPG = file.raw.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      if(isJPG && isLt2M){
        this.imageUrl = window.URL.createObjectURL(file.raw);
      }
      return isJPG && isLt2M;
    }
  }
};
</script>