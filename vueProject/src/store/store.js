import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

var mutations = {
    setUser(state,user){
        state.user = user
        sessionStorage.user = JSON.stringify(user)
    },
    setUserAvatar(state,url){
        state.user.avatar = url
        var user = state.user
        sessionStorage.user = JSON.stringify(user)
    },
    setUserNickname(state,nickname){
        state.user.nickname = nickname
        var user = state.user
        sessionStorage.user = JSON.stringify(user)
    },
    setUserDescription(state,description){
        state.user.description = description
        var user = state.user
        sessionStorage.user = JSON.stringify(user)
    },
    delUser(state) {
        state.token = {}
        sessionStorage.removeItem('user')
    }
}


var store = new Vuex.Store({
    state :{
        user : {},
    },
    mutations
})

export default store