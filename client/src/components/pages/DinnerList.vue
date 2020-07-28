<template>

  <div>
    <v-card>
      <v-card-title>ツイート</v-card-title>
      <div v-if="dinnerErrors">
        <v-alert
          dense
          type="error"
        >
          <ul>
            <li v-if="dinnerErrors.text">text：{{ dinnerErrors.text }}</li>
            <li v-if="dinnerErrors.category_id">category_id：{{ dinnerErrors.category_id }}</li>
          </ul>
        </v-alert>
      </div>
      <v-card-text>
        <v-text-field
          v-model="dinnerText"
          label="text"
        >
        </v-text-field>
        <v-text-field
          v-model="dinnerCategory_id"
          label="category_id"
        >
        </v-text-field>
        <v-card-actions>
          <v-btn @click="createDinner">つぶやく</v-btn>
        </v-card-actions>
      </v-card-text>
      <v-divider></v-divider>
    </v-card>
    <ul>
      <li v-for="dinner in dinners" :key="dinner.id">
        ID；{{ dinner.ID }}、
        Text：{{ dinner.text }}、
        Category：{{ category_label[dinner.category_id] }}、
        <RouterLink
          :to="`/dinner/${dinner.ID}`"
        >|詳細へ|</RouterLink>
        <RouterLink
          :to="`/dinner/delete/${dinner.ID}/confirm`"
        >|削除|</RouterLink>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data () {
    return {
      dinners: [],
      dinnerErrors: '',
      dinnerText: '',
      dinnerCategory_id: '',
      category_label: {
        1 : "デート",
        2 : "パーティー",
        3 : "家族ごはん",
        4 : "穴場",
        5 : "テイクアウト",
      }
    }
  },
  methods: {
    async getAllDinners () {
      const response = await this.$axios.get(`/api/dinner`,
        { headers: { 'Authorization': `Bearer ${localStorage.getItem('jwt')}` }
      })
        .catch(err => err.response || err)
        console.log(response)
      this.dinners = response.data
    },
    async createDinner () {
      const params = new URLSearchParams();
      params.append('text', this.dinnerText)
      params.append('category_id', this.dinnerCategory_id)

      const response = await this.$axios.post(`/api/dinner/create`, params,
        { headers: { 'Authorization': `Bearer ${localStorage.getItem('jwt')}` }})
        .catch(err => err.response || err)

        if (response.status === 400) {
          this.dinnerErrors = response.data.error
        } else {
          this.dinnerErrors = ''
          this.dinnerText = '',
          this.dinnerCategory_id = '',
          this.getAllDinners()
        }
    },
  },
  watch: {
    $route: {
      async handler () {
        await this.getAllDinners()
      },
      immediate: true
    }
  }
}
</script>