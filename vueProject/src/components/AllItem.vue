<template>
  <div style="margin-top:30px">
    <el-row :gutter="20" v-for="count in lines" :key="count" style="margin-top:20px">
      <el-col
        :span="4"
        v-for="(v,k) in d.slice((count-1) * 6,(count == 1?6:d.length))"
        :key="k"
        
      >
        <el-link :underline="false" @click.native="onClick(v.uuid)">
          <el-image style="width:220px;height:220px;" :src="v.avatar" fit="fill"></el-image>
          <p style="margin:0px;text-align:left">{{ v.title }}</p>
        </el-link>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      d: [],
      lines: 0,
      name: this.$store.state.token,
      nil: "" == this.$store.state.token
    };
  },
  methods: {
    onClick(uuid) {
      this.$router.push({ path: "/main/item/" + uuid });
    }
  },
  mounted() {
    //axios.defaults.withCredentials = true
    this.axios.get("http://localhost:8000/api/v1/video/get").then(res => {
      var data = res.data.data;
      this.lines = Math.floor(data.length / 6) + 1;

      for (var i = 0; i < data.length; i++) {
        data[i].avatar =
          "http://localhost:8000" +
          data[i].avatar.substring(1, data[i].avatar.length);
      }
      this.d = data;
    });
  }
};
</script>

<style>
</style>