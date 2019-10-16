<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
import "./assets/css/iconfont.css";
export default {
  name: "App",
  methods: {
    saveUser() {
      sessionStorage.setItem("user", JSON.stringify(this.$store.state.user));
      console.log(sessionStorage.getItem("user"));
    }
  },
  created() {
    window.addEventListener("unload", this.saveUser);
    this.$store.commit("setUser", JSON.parse(sessionStorage.user));
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}
</style>
