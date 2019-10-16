<template>
  <div style="margin-top:80px">
    <el-row>
      <el-col :span="12" :offset="6">
        <el-row>
          <el-col :span="24">
            <el-form label-width="80px" :model="form">
              <el-form-item label="视频">
                <el-col :span="14">
                  <el-upload
                    ref="videoUpload"
                    drag
                    :auto-upload="false"
                    style="text-align:left;line-height:normal"
                    action="/"
                    :limit="1"
                    v-loading="videoProgress"
                  >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">
                      将文件拖到此处，或
                      <em>点击上传</em>
                    </div>
                  </el-upload>
                </el-col>
                <el-col :span="5">
                  <el-col :span="24" v-if="videoProgress">
                    <el-col :span="6" style="line-height:normal;">
                      <el-progress
                        type="circle"
                        :percentage="videoMd5Percentage"
                        :status="Md5Stauts"
                        :width="20"
                        :stroke-width="2"
                        :show-text="false"
                      ></el-progress>
                    </el-col>
                    <el-col :span="14" style="line-height:normal;">md5校验</el-col>
                  </el-col>
                  <el-col :span="24" style="margin-top:50px">
                    <el-progress
                      :text-inside="true"
                      :stroke-width="26"
                      :percentage="videoUploadPercentage"
                      v-if="videoProgress"
                    ></el-progress>
                  </el-col>
                </el-col>
              </el-form-item>
              <el-form-item label="封面">
                <el-col :span="5">
                  <el-upload
                    class="avatar-uploader"
                    action="/api/api/v1/image/upload"
                    ref="avatarUpload"
                    :show-file-list="false"
                    :on-success="handleAvatarSuccess"
                    :on-change="onAvatarChange"
                    :auto-upload="false"
                  >
                    <img v-if="imageUrl" :src="imageUrl" class="avatar" />
                    <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                  </el-upload>
                </el-col>
              </el-form-item>
              <el-form-item label="标题">
                <el-input v-model="form.title"></el-input>
              </el-form-item>
              <el-form-item label="描述">
                <el-input v-model="form.description"></el-input>
              </el-form-item>
            </el-form>
            <el-button type="button" @click="onClick">提交</el-button>
          </el-col>
        </el-row>
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
import sparkMD5 from "spark-md5";
import { Loading, Main } from "element-ui";
import { resolve } from 'url';
import { reject } from 'q';
import router from '../router';

export default {
  data() {
    return {
      percentage: 0,
      imageUrl: "",
      Md5Stauts: "success",
      videoProgress: false,
      videoUploadPercentage: 0,
      videoMd5Percentage: 0,
      form: {
        title: "",
        description: "",
        md5: "",
        videoUuid: "",
        avatarUrl: ""
      }
    };
  },
  watch: {
    Md5Stauts(val, oldVal) {
      if (val == 100) {
        Md5Stauts = "success";
      }
    }
  },
  methods: {
    onClick1() {
      this.percentage += 10;
      return;
    },
    calcMd5(file) {
      var that = this
      return new Promise((resolve,reject) => {
        var fr = new FileReader()
        var spark = new sparkMD5()
        fr.onload = (e) => {
          spark.append(e.target.result)
          resolve(spark.end())
        }
        fr.onprogress = (e) => {
          that.videoMd5Percentage = Math.floor((e.loaded / e.total) * 100)
        }
        fr.readAsBinaryString(file.raw)
      })
    },

    submitChunk(file, md5) {
      var that = this
      return new Promise((resolve,reject)=>{
        var chunkSize = 2 * 1024 * 1024,
        chunks = Math.ceil(file.size / chunkSize),
        currentChunk = 0;
        var result = {}
        while (currentChunk < chunks) {
          that.videoUploadPercentage = Math.floor((currentChunk/chunks) * 100)
          var start = currentChunk * chunkSize,
            end = start + chunkSize > file.size ? file.size : start + chunkSize;
          var currentFile = file.raw.slice(start, end);
          
          var f = new FormData();
          f.append("file", currentFile);
          f.append("currentChunk", currentChunk);
          f.append("totalChunks", chunks);
          f.append("uuid",this.form.videoUuid)  
          f.append("md5", md5);
          f.append("fileType",file.raw.type.split("/")[1])
          result = this.axios.post(this.BaseUrl + "/api/v1/video/upload", f)
          currentChunk ++
        }
        resolve(result)
      })
    },

    async onClick() {
      this.videoProgress = true;
      var file = this.$refs.videoUpload.uploadFiles[0];
      var file1 = this.$refs.avatarUpload.uploadFiles[0];
      // var md5 = await this.calcMd5(file)
      

      this.axios.post(this.BaseUrl + "/api/v1/videoform/upload",this.form).then((res)=>{
        this.form.videoUuid = res.data.data
        return this.calcMd5(file)
      }).then((res) => {
        this.form.md5 = res
        return this.axios.post(this.BaseUrl + "/api/v1/videoform/checkmd5/" + res + "?uuid=" + this.form.videoUuid )
      }).then((res) => {
        res = res.data
        if(res.data == 0){
          return this.submitChunk(file,this.form.md5)
        }else{
          return
        }
      }).then((res) => {
        this.videoUploadPercentage = 100;
        var f = new FormData()
        f.append("videoUuid",this.form.videoUuid)
        f.append("file",file1.raw)
        return this.axios.post(this.BaseUrl + "/api/v1/image/video/upload",f)
      }).then((res)=>{
        if(res.data.status == 50000){
          alert("上传成功!")
          this.$router.push("/main")
        }
      })


      // this.calcMd5(file).then((res) => {
      //   this.form.md5 = res
      //   return this.submitChunk(file,res)
      // }).then((res) => {
      //   this.form.videoUuid = res.msg
      //   this.$refs.avatarUpload.submit()
      //   return this.axios.post("/api/api/v1/videoform/upload",this.form)
      // }).then((res)=>{
      //   console.log(res)
      // }).catch((e) => {
      //   console.log(e)
      // })

      // console.log("@######")

      // if (!file) {
      //   this.$message.error("请选择视频！");
      //   return;
      // }

      // if (!file1) {
      //   this.$message.error("请选择头像！");
      //   return;
      // }
    },

    handleAvatarSuccess(res, file) {
      this.form.avatarUrl = res.msg
    },
    onAvatarChange(file) {
      file = file.raw;
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      this.imageUrl = URL.createObjectURL(file);
      return isJPG && isLt2M;
    }
  }
};
</script>