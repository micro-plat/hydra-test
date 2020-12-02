<template>
  <div class="hello">
      <button @click="clickquery">Options</button>
      <label>{{requestData}}</label>
  </div>
</template>

<script>

import axios from "axios"

export default {
  name: 'Options',
  props: {
    msg: String
  },
  data(){
    return {
      requestData:""
    }
  },
  methods:{
    clickquery(){
    var url= window.globalconfig.api_addr+"/options"
    var that = this;
    axios.get(url).then(function(res){
      console.log(res);
      that.requestData = res
    }).catch(function (error) {
        console.log(error);
    });
    }
  }
}
</script>
 