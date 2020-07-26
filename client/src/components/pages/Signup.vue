<template>
  <div>
    <v-card class="elevation-1 pa-3 login-card">
      <v-card-text>
        <div class="layout column align-center">
          <h1 class="flex my-4 primary--text font-weight-bold">新規登録</h1>
        </div>
        <div v-if="signUpErrors">
          <v-alert
            dense
            type="error"
          >
            <ul>
              <li>{{ signUpErrors.error }}</li>
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
          <v-btn block color="primary" @click="signUp">新規登録</v-btn>
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
      signUpErrors: ''
    }
  },
  methods: {
    async signUp () {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
        .then(res => {
          this.$router.push('/login')
        }).catch(error => {
          console.log(error.message)
        })
      // const params = new URLSearchParams();
      // params.append('email', this.email)
      // params.append('password', this.password)
      // console.log(params)

      // const response = await this.$axios.post(`/api/signup`, params)
      //   .catch(err => err.response || err)
      //   console.log(response.data.error)
      //   if (response.status === 400) {
      //     this.signUpErrors = response.data
      //   } else {
      //     this.$router.replace({ path: '/login' })
      //   }
    },
  }
}
</script>