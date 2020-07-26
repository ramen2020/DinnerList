<template>
  <div>わたしのなまえは、{{name}}
    <h1 class="font-weight-light">{{ $route.meta.name }}</h1>
    <p>ID：{{ dinner.ID }}</p>
    <p>Text：{{ dinner.text }}</p>
    <p>カテゴリ：{{ dinner.category_id }}</p>
    <p>作成日：{{ dinner.CreatedAt }}</p>
    <v-btn>
      <RouterLink
        :to="`/dinner/delete/${dinner.ID}/confirm`"
      >削除</RouterLink>
    </v-btn>
  </div>
</template>

<script>
import firebase from 'firebase'
export default {
  props: {
    id: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      dinner: null,
      name: firebase.auth().currentUser.email
    }
  },
  methods: {
    async getDinnerById() {
      const response = await this.$axios.get(`/api/dinner/${this.id}`)
      console.log(response.data)
      this.dinner = response.data
    },
  },
  watch: {
    $route: {
      async handler () {
        await this.getDinnerById()
      },
      immediate: true
    }
  }
}
</script>