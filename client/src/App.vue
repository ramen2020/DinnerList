<template>
  <v-app>
    <Navbar />
    <v-content>
      <RouterView />
    </v-content>
    <Footer />
  </v-app>
</template>

<script>
import Navbar from './components/Navbar.vue'
import Footer from './components/Footer.vue'

export default {
  components: {
    Navbar,
    Footer
  }
}

</script>

<style>

</style>
