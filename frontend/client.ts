import '@babel/polyfill'
import 'vuetify/dist/vuetify.min.css'
import '@mdi/font/css/materialdesignicons.css'

import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false;

new Vue({
  template: '<App/>',
  components: {App},
  render: h => h(App)
}).$mount('#app');
