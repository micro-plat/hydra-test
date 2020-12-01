import Vue from 'vue'
import App from './App.vue'
import axios from "axios"

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

newFunction();

function newFunction() {
  axios.get('/vue/config').then(function(res){
    window.globalconfig=res.data
    console.log("main.js:",window.globalconfig);
}).catch(function (error) {
    console.log(error);
});
}
