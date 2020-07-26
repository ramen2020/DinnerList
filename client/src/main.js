import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import vuetify from './plugins/vuetify';
import firebase from 'firebase'

Vue.prototype.$axios = axios
Vue.config.productionTip = false

const config = {
  apiKey: process.env.VUE_APP_apiKey,
  authDomain: process.env.VUE_APP_authDomain,
  databaseURL: process.env.VUE_APP_databaseURL,
  projectId: process.env.VUE_APP_projectId,
  storageBucket: process.env.VUE_APP_storageBucket,
  messagingSenderId: process.env.VUE_APP_messagingSenderId,
  appId: process.env.VUE_APP_appId,
  measurementId: process.env.VUE_APP_measurementId
}

firebase.initializeApp(config)

let app
firebase.auth().onAuthStateChanged(user => {
  /* eslint-disable no-unused-vars */
  if (!app) {
    new Vue({
      router,
      store,
      vuetify,
      firebase,
      render: h => h(App)
    }).$mount('#app')
  }
/* eslint-enable no-unused-vars */
})
