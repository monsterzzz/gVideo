<template>
  <el-row class="tac">
   
  <el-col :span="4">
    <h5>个人信息</h5>
    <el-menu
      :default-active="$route.path"
      class="el-menu-vertical-demo"
      :router="true"
      @open="handleOpen"
      @close="handleClose">
      <el-menu-item index="/main/user/info">
        <i class="el-icon-menu"></i>
        <span slot="title">导航1</span>
      </el-menu-item>
      <el-menu-item index="/main/user/avatar">
        <i class="el-icon-document"></i>
        <span slot="title">导航2</span>
      </el-menu-item>
      <el-menu-item index="/main/user/videorecord" >
        <i class="el-icon-setting"></i>
        <span slot="title">导航3</span>
      </el-menu-item>
    </el-menu>
  </el-col>
  <el-col :span="20">
    <router-view></router-view>
  </el-col>
</el-row>

</template>
<script>
  export default {
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>