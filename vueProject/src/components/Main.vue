<template style="margin-top:0px">
  <div>
    <el-row>
      <el-col :span="24">
        <el-menu router :default-active="$route.matched[1].path" mode="horizontal" @select="handleSelect">
          <!-- <el-menu-item index="1">处理中心</el-menu-item>
            <el-menu-item index="2">处理中心</el-menu-item>-->

          <el-menu-item
            v-for="(value,key) in f"
            :key="key"
            :index="value.path"
          >{{ key }} - {{ value.name }}</el-menu-item>
          <!-- <el-menu-item index="3" style="float:right">处理中心1</el-menu-item> -->
          <!-- <el-menu-item index="1" >处理中心1</el-menu-item>
          <el-menu-item index="2" >处理中心1</el-menu-item>

          <el-menu-item index="3" >处理中心1</el-menu-item>

          <el-menu-item index="4" >处理中心1</el-menu-item>
          <el-menu-item index="5" >处理中心1</el-menu-item>
          <el-menu-item index="6" >处理中心1</el-menu-item> -->


        </el-menu>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="20" :offset="2">
        <router-view></router-view>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Upload from "@/components/Upload";
import User from "@/components/User";

export default {
  data() {
    return {
      f: [
        {
          name: "all",
          path: "/main/content"
        },
        {
          name: "upload",
          path: "/main/upload"
        },
        {
          name: "user",
          path: "/main/user"
        }
      ]
    };
  },
  
  methods: {
    handleSelect(key, keyPath) {
      console.log(this.$route.matched[1].path)
    }
  }
  //   components : {
  //       Upload,
  //       User
  //   }
};
</script>

<style>
</style>