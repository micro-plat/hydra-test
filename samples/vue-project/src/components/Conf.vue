<template>
  <div class="hello"> 
    从服务器读取vue 配置信息:{{config}}
  </div>
</template>

<script>
export default {
  name: 'Conf',
  props: {
    msg: String
  },
  data(){
    var config = JSON.stringify( window.globalconfig)
    return {
      config: config
    }
  },  
}
</script>
 