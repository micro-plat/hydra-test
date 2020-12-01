<template>
  <div class="hello">
      this is web page 
      <button @click="queryapi">点击验证接口请求</button>
  </div>
</template>

<script>

import axios from "axios"

export default {
  name: 'Mixed',
  props: {
    msg: String
  },
  methods:{
    queryapi(){
      axios.get('/self/api').then(function(res){
          console.log("apiresponse:",res.data);
      }).catch(function (error) {
          console.log(error);
      });
    }
  }
}
</script>
 