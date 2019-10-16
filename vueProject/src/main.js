// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import axios from 'axios'
import VueAxios from 'vue-axios'
import store from './store/store.js'



window.Hls = require("hls.js")

axios.defaults.withCredentials = true

Vue.use(Element)
Vue.use(VueAxios,axios)
Vue.config.productionTip = false
Vue.prototype.BaseUrl = "http://localhost:8000"



router.beforeEach((to, from, next)=>{
  //if(to.)
  next()
})


/* eslint-disable no-new */
var s = new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})

