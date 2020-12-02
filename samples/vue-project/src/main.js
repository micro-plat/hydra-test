import Vue from 'vue'
import App from './App.vue'
import axios from "axios"

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

initAxiosConfig();

loadServerConfig();

function initAxiosConfig(){
  axios.defaults.timeout = 10000
  axios.defaults.baseURL = '/'
  axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
  axios.defaults.withCredentials = true
}



function loadServerConfig() {
  axios.get('/vue/config').then(function(res){
    window.globalconfig=res.data
    console.log("main.js:",window.globalconfig);
}).catch(function (error) {
    console.log(error);
});
}
