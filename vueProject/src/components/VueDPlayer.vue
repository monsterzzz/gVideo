<template>
    <div id="dplayer">
        
    </div>
</template>

<script>
require("../../node_modules/dplayer/dist/DPlayer.min.css")
import DPlayer from 'DPlayer'

export default {
    props:{
        size:{
            type:String,
            default:"16:9"
        },
        options:{
            type:Object,
            default:{
                video :{
                    url:""
                }
            }
        }
    },
    data() {
        
        return {
            dp:null
        }
    },
    mounted(){
        var size = this.size.split(":")
        var height = (this.$el.offsetWidth * parseInt(size[1])) /  parseInt(size[0])
        console.log("111",height)
        this.$el.style.height = height + "px"
        this.options.element = this.$el
        
       
        const player = this.dp = new DPlayer(this.options)
        console.log(this.options.video.url)
        console.log("!!!")

        player.on("play",()=>{
            console.log("!!")
            this.$emit("play")
        })
    },

}
</script>