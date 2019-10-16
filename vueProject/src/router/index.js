import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Item from '@/components/Item'
import SignIn from '@/components/SignIn'
import Main from '@/components/Main'
import Upload from '@/components/Upload'
import User from '@/components/User'
import AllItem from '@/components/AllItem'
import Info from '@/components/Info'
import Avatar from '@/components/Avatar'
import VideoRecord from '@/components/VideoRecord'



Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    
    {
      path: "/signIn",
      name : "SignIn",
      component : SignIn
    },
    {
      path: "/main",
      name : "Main",
      redirect : "main/content",
      component : Main,
      children:[
        {
          path: '/main/item/:uuid',
          name: 'Item',
          component: Item
        },
        {
          path : "/main/content",
          component : AllItem
        },
        {
          path : "/main/upload",
          component : Upload
        },
        {
          path : "/main/user",
          name:"User",
          redirect : "/main/user/info",
          component : User,
          children:[
            {
              path : "/main/user/info",
              component : Info
            },
            {
              path : "/main/user/avatar",
              component : Avatar
            },
            {
              path : "/main/user/videorecord",
              component : VideoRecord
            },
            
          ]
        },
        
      ]
    },
  ]
})
