<template>
<div>
  <v-app-bar>
    <RouterLink to="/">
      <v-btn>twi</v-btn>
    </RouterLink>
    <v-spacer></v-spacer>
    <div v-if="authenticatedUser">
      <v-btn @click=signOut>Logout</v-btn>
    </div>
    <div v-else>
      <RouterLink to="/login">
        <v-btn>Login</v-btn>
      </RouterLink>
      <RouterLink to="/signup">
        <v-btn>SignUp</v-btn>
      </RouterLink>
    </div>
  </v-app-bar>
</div>
</template>

<script>
import firebase from 'firebase'
export default {
  data () {
    return {
      authenticatedUser: false,
    }
  },
  methods: {
    signOut: function () {
      firebase.auth().signOut().then(() => {
        localStorage.removeItem('jwt')
        this.$router.push('/login')
        })
    }
  },
  mounted(){
    firebase.auth().onAuthStateChanged((user) => {
      if (user) {
        console.log('login');
        this.authenticatedUser = true;
      } else {
        console.log('logout');
        this.authenticatedUser = false;
      }
    });
  }
}
</script>