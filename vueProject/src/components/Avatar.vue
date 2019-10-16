<template>
  <div style="margin-top:3em">
    <el-row>
      <el-col :span="15" :offset="4">
        <el-card>
            <p style="text-align:left;margin-left:20px;margin-top:0px">头像</p>
            <el-divider><el-avatar :size="70" :src="user.avatar"></el-avatar></el-divider>
          <el-col :span="20" :offset="1" style="margin-top:3.5em">
              <el-col :span="21">
                <el-row>
                  <el-col
                    :span="10"
                    :offset="5"
                    style="border-right-style:dashed;border-right-width:1px;"
                  >
                    <el-upload
                      ref="avatar"
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

                  <el-col :span="3" :offset="3" style="text-align:left">
                    <el-avatar :size="70" :src="imageUrl" style="margin-top:3.5em"></el-avatar>
                  </el-col>
                </el-row>
                <el-row>
                  <el-col :span="6" :offset="12" style="margin-top:20px;padding-bottom:30px">
                    <el-button @click="submit">修改</el-button>
                  </el-col>
                </el-row>
              </el-col>
            
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
      imageUrl: "",

      user: {
        username: "",
        nickname: "",
        description: "",
        avatar: ""
      }
    };
  },

  methods: {
    handleAvatarSuccess(res, file) {
      console.log(res);
      this.$message.info("上传成功!");
    },
    submit() {
      var f = new FormData();
      console.log(this.$refs.avatar.uploadFiles);
      f.append("file", this.$refs.avatar.uploadFiles[0].raw);
      this.axios
        .post(this.BaseUrl + "/api/v1/image/avatar/upload", f)
        .then(res => {
          var data = res.data;
          if (data.status == 50000) {
            console.log(data.data);
            this.$store.commit(
              "setUserAvatar",
              this.BaseUrl + "" + data.data.path
            );
          }
          console.log(this.$store.state);
          this.user = this.$store.state.user
        });
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
      if (isJPG && isLt2M) {
        this.imageUrl = window.URL.createObjectURL(file.raw);
      }
      return isJPG && isLt2M;
    }
  },
  mounted() {
    var user = this.$store.state.user;
    this.user = user;
    console.log(user);
    console.log(this.user);
  }
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