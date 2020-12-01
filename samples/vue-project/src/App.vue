<template>
  <div id="app">
    <Conf msg="Welcome to Your Vue.js App"/>
    <Mixed msg="Welcome to Your Vue.js App"/>
    <Options msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import Conf from "./components/Conf.vue";
import Mixed from "./components/Mixed.vue";
import Options from "./components/Options.vue";

import axios from "axios"

export default {
  name: "App",
  components: {
    Conf,
    Mixed,
    Options
  },
  mounted(){
    axios.get('/vue/config').then(function(res){
        window.globalconfig=res.data
        console.log("globalconfig:",window.globalconfig);
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
