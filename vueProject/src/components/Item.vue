<template>
  <div style="margin-top:1em">
    <el-row>
      <el-col :span="18">
        <el-row>
          <el-row>
            <el-col :span="24">
              <p
                style="margin:0px;font-size:18px;margin-bottom:5px;text-align:left;"
              >{{ videoInfo.title ? videoInfo.title : "error" }}</p>
              <p
                style="margin:0px;font-size:10px;margin-bottom:5px;text-align:left;"
              >{{ videoInfo.time ? videoInfo.time : "error" }}</p>
              <el-card :body-style="{padding:'0px'}">
                <d-player @play="play" ref="player" :options="options" :size="size">1221</d-player>
                <el-row style="margin-top:10px;margin-bottom:5px">
                  <el-col
                    :span="4"
                    style="display: table-cell;vertical-align: middle;text-align: center;height:100% "
                  >
                    <div style="margin-top:0.5em;margin-bottom:0px">
                      <div style="display:inline;margin-left:5px">
                        <i style="font-size:20px;" class="icon-92"></i>
                        <span style="margin-left:10px">{{videoInfo.play_count}}</span>
                      </div>
                      <div style="display:inline;margin-left:10px">
                        <i style="font-size:20px;" class="icon-100"></i>
                        <span style="margin-left:10px">{{videoInfo.likes_count}}</span>
                      </div>
                    </div>
                  </el-col>
                  <el-col :span="15" :offset="1">
                    <el-input v-model="danmu_ipt" placeholder="请输入内容"></el-input>
                  </el-col>
                  <el-col :span="3">
                    <el-button type="primary" @click="draw">主要按钮</el-button>
                  </el-col>
                </el-row>
              </el-card>
            </el-col>
          </el-row>
          <el-row style="margin-top:40px">
            <el-col :span="24">
              <el-card class="box-card" :body-style="{padding:'5px'}">
                <div slot="header" class="clearfix" style="text-align:left">
                  <span>评论区</span>
                </div>
                <el-row
                  style="padding-bottom:30px;padding-top:30px;border-bottom: 1px solid #EBEEF5;"
                >
                  <el-col :span="1" :offset="1" style="text-align:left;margin-right:10px">
                    <!-- <p style="margin-bottom:0px"></p> -->
                    <el-avatar :size="50" :src="user.avatar"></el-avatar>
                  </el-col>
                  <el-col :span="19">
                    <el-input type="textarea" :rows="3" placeholder="请输入内容" v-model="textarea"></el-input>
                  </el-col>
                  <el-col :span="2" style="margin-left:10px">
                    <p>
                      <el-button type="primary" @click="submitComment">发表</el-button>
                    </p>
                  </el-col>
                </el-row>
                <div style="margin-top:20px">
                  <div v-for="(o,i) in comments" :key="i" class="text item">
                    <el-card class="box-card" shadow="never" style="margin-bottom:5px">
                      <div style="text-align:left;">
                        <el-row style="">
                          <el-col :span="2">
                            <el-avatar :size="50" :src="'http://localhost:8000' + o.user.avatar"></el-avatar>
                          </el-col>
                          <el-col :span="20">
                            <el-row style="margin-bottom:10px">
                              <el-col :span="12" style="font-size:12px;font-weight:700">{{ o.user.nickname }}</el-col>
                            </el-row>
                            <el-row style="margin-bottom:10px">
                              <el-col :span="24" style="font-size:14px;">{{ o.comment.content }}</el-col>
                            </el-row>
                            <el-row>
                              <el-col :span="24" style="font-size:12px;">{{ o.comment.time }}</el-col>
                            </el-row>
                          </el-col>
                        </el-row>
                      </div>
                    </el-card>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
        </el-row>
      </el-col>
      <el-col :span="5" style="text-align:left;margin-left:30px">
        <el-row>
          <el-col :span="24">
            <el-card class="box-card">
              <el-divider>
                <el-avatar :size="50" :src="videoInfo.user.avatar"></el-avatar>
              </el-divider>
              <div style="margin-top:2em;text-align:center">
                <p>{{ videoInfo.user.nickname ? videoInfo.user.nickname : "该用户很懒,没有任何描述" }}</p>
                <p>{{ videoInfo.user.description ? videoInfo.user.description : "该用户很懒,没有任何描述" }}</p>
              </div>
            </el-card>
          </el-col>
        </el-row>
        <el-row style="margin-top:30px">
          <el-col :span="24">
            <el-card class="box-card" :body-style="{padding:'5px'}">
              <div slot="header" class="clearfix">推荐列表</div>
              <div v-for="o in 7" :key="o" class="text item">
                <el-card class="box-card" shadow="never" style="margin-bottom:5px">
                  <div class="text item">{{'列表内容 ' + "o" }}</div>
                </el-card>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import VueDplayer from "./VueDPlayer";
import dayjs from "dayjs"

export default {
  name: "sss",
  data() {
    return {
      user: {},
      owner: {},
      videoInfo: {
        user: {}
      },
      danmu_ipt: "",
      textarea: "",
      comments:{

      },
      size: "100:65",
      options: {
        video: {
          url: "",
          type: "hls"
        },
        danmaku: {
          id: "demo",
          api: "https://api.prprpr.me/dplayer/"
        }
      }
    };
  },
  methods: {
    play() {
      console.log("play");
    },
    draw() {
      console.log(this.$refs.player);
      console.log(this.$refs.player.dp);
      this.$refs.player.dp.danmaku.draw({
        text: this.danmu_ipt,
        color: "#b7daff",
        type: "right" // should be `top` `bottom` or `right`
      });
      this.danmu_ipt = "";
    },
    submitComment() {
      this.axios
        .post(this.BaseUrl + "/api/v1/comment/add", {
          video_uuid: this.$route.params.uuid,
          content: this.textarea
        })
        .then(res => {
          
          if(res.data.data == "success"){
            var newComment = {
              user :{
                nickname:this.user.nickname,
                avatar:this.user.avatar.replace(this.BaseUrl + "",""),
              },
              comment:{
                content:this.textarea,
                time:dayjs().format("YYYY-MM-DD HH:mm:ss")
              }
            }
            this.$message.success("评论成功~")
            this.textarea = ""
            this.comments.unshift(newComment)
            
          }
        });
    }
  },
  computed: {},
  created() {},
  components: {
    "d-player": VueDplayer
  },
  mounted() {
    console.log("mounted");
    
    console.log(this.$el);
    this.user = this.$store.state.user;
    this.player = this.$refs.player.dp;
    this.player.video.setAttribute("style", "object-fit:fill");

    this.axios.get(this.BaseUrl + "/api/v1/comment/" + this.$route.params.uuid ).then(res=>{
      var data = res.data.data
      data.sort(function(a,b){
        if(dayjs(a.comment.time) > dayjs(b.comment.time)){
          return -1
        }else{
          return 1
        }

      })
      this.comments = data

    })

    this.axios
      .get(this.BaseUrl + "/api/v1/video/item/" + this.$route.params.uuid)
      .then(res => {
        var data = res.data.data;
        // console.log(data)
        data.user.avatar = this.BaseUrl + "" + data.user.avatar;
        this.videoInfo = data;
        //this.options.video.url = data.m3u8;
        this.player.switchVideo({ url: this.BaseUrl + "" + data.m3u8 });
      });
    
  }

  
};
</script>
<style scoped>
.fill {
  object-fit: cover;
}
</style>