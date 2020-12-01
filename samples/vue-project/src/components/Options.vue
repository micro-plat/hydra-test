<template>
  <div class="hello">
      <button @click="clickquery">Options</button>
  </div>
</template>

<script>

import axios from "axios"

export default {
  name: 'Options',
  props: {
    msg: String
  },
  methods:{
    clickquery(){
    var url= window.globalconfig.api_addr+"/options"
    axios.get(url).then(function(res){
      console.log(res);
    }).catch(function (error) {
        console.log(error);
    });
    }
  }
}
</script>
 