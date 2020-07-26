<template>
  <div>
    <v-card class="elevation-1 pa-3 login-card">
      <v-card-text>
        <div class="layout column align-center">
          <h1 class="flex my-4 primary--text font-weight-bold">ログイン</h1>
        </div>
        <div v-if="loginErrors">
          <v-alert
            dense
            type="error"
          >
            <ul>
              <li>{{ loginErrors.error }}</li>
            </ul>
          </v-alert>
        </div>
        <v-form>
          <v-text-field
            v-model="email"
            label="email"
          >
          </v-text-field>
          <v-text-field
            v-model="password"
            label="password"
          >
          </v-text-field>
        </v-form>
        <div class="login-btn">
          <v-btn block color="primary" @click="login">ログイン</v-btn>
        </div>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
import firebase from 'firebase'
export default {
  data () {
    return {
      email: '',
      password: '',
      loginErrors: ''
    }
  },
  methods: {
    async login () {
      const response = firebase.auth().signInWithEmailAndPassword(this.email, this.password)
        .then(res => {
        res.user.getIdToken().then(idToken => {
        localStorage.setItem('jwt', idToken.toString())
    })
        }, err => {
          alert(err.message)
        })

      // const params = new URLSearchParams();
      // params.append('email', this.email)
      // params.append('password', this.password)
      // console.log(params)

      // const response = await this.$axios.post(`/api/login`, params)
      //   .catch(err => err.response || err)
      //   console.log(response.data.error)
      //   if (response.status === 400) {
      //     this.loginErrors = response.data
      //   } else {
      //     this.$router.replace({path: '/'})
      //   }
    },
  }
}
</script>