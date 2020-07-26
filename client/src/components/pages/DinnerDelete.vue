<template>
  <div>
    <h1 class="font-weight-light">{{ $route.meta.name }}</h1>
    <h1>削除しますか？</h1>
    <p>ID：{{ dinner.ID }}</p>
    <p>Text：{{ dinner.text }}</p>
    <p>カテゴリ：{{ dinner.category_id }}</p>
    <p>作成日：{{ dinner.CreatedAt }}</p>
    <v-btn @click=getDeleteById>削除！！！</v-btn>
  </div>
</template>
<script>
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
    }
  },
  methods: {
    async getDeleteById() {
      await this.$axios.post(`/api/dinner/delete/${this.id}`)
            this.$router.replace({path: '/'})
    },
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