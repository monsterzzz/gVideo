


<template>
  <div style="margin-top:3.5em">
    <el-row>
      <el-col :span="20" :offset="2" v-loading="loading">
        <el-card class="box-card">
          <!-- <div slot="header" style="text-align:left">
            <span>投稿纪录</span>
            <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
          </div>-->
          <div>
            <el-table :data="tableData" style="width: 100%">
              <el-table-column prop="time" label="日期" width="180"></el-table-column>
              <el-table-column prop="picture" label="封面" width="100">
                <template slot-scope="scope">
                  <el-image :src="scope.row.picture" style="width:50px" ></el-image>
                </template>
              </el-table-column>
              <el-table-column prop="title" :show-overflow-tooltip="true" label="标题"></el-table-column>
              <el-table-column prop="description" :show-overflow-tooltip="true" label="描述"></el-table-column>
              <el-table-column prop="play_count" label="播放数"></el-table-column>
              <el-table-column prop="likes_count" label="播放数"></el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

  <script>
export default {
  data() {

    return {
        loading:true,
      tableData: []
    };
  },
  mounted() {
    this.axios
      .get(this.BaseUrl + "/api/v1/user/videorecords")
      .then(res => {
        var data = res.data.data;
        for(var i=0;i<data.length;i++){
            data[i].picture = this.BaseUrl + "" + data[i].picture.substring(1)
        }
        this.tableData = data;
        this.loading = false
      });
  }
};
</script>