<template>
  <div id="app"> 
      <component v-bind:is="curComponent"></component>
  </div>
</template>

<script>
import Conf from "./components/Conf.vue";
import Mixed from "./components/Mixed.vue";
import Options from "./components/Options.vue";
import Static from "./components/Static.vue";

import axios from "axios"

export default {
  name: "App",
  data(){
    return {
       curComponent:null,
    }
  },
  components: {
    Conf,
    Mixed,
    Options,
    Static
  },
  mounted(){
    var componentsMap = {
      "conf":Conf,
      "mixed":Mixed,
      "options":Options,
      "static":Static,
    } 
    console.log(componentsMap)
    var that = this;
    axios.get('/vue/config').then(function(res){
        window.globalconfig=res.data
        console.log("globalconfig:",window.globalconfig);
        that.curComponent =componentsMap[window.globalconfig.currentComponent]
    }).catch(function (error) {
        console.log(error);
    });
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
