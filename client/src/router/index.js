import Vue from 'vue'
import VueRouter from 'vue-router'

// ページコンポーネント
import Login from '../components/pages/Login.vue'
import SignUp from '../components/pages/Signup.vue'
import DinnerList from '../components/pages/DinnerList.vue'
import DinnerDetail from '../components/pages/DinnerDetail.vue'
import DinnerDelete from '../components/pages/DinnerDelete.vue'

import firebase from 'firebase'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    name: 'top',
    component: DinnerList,
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    component: Login,
  },
  {
      path: '/signup',
      component: SignUp
    },
  {
    path: '/dinner/:id',
    component: DinnerDetail,
    props: true,
    meta: { requiresAuth: true }
    },
  {
    path: '/dinner/delete/:id/confirm',
    component: DinnerDelete,
    props: true,
    meta: { requiresAuth: true }
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  let currentUser = firebase.auth().currentUser
  let requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  if (requiresAuth && !currentUser) next('/login')
  else if (!requiresAuth && currentUser) next()
  else next()
})


export default router
