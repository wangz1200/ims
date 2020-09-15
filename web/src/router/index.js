import Vue from 'vue'
import VueRouter from 'vue-router'

import LoanQuery from "../views/LoanQuery"

Vue.use(VueRouter)

  const routes = [
  {
    path: "/loan/query",
    name: 'LoanQuery',
    component: LoanQuery,
  },
  {
    path: '/loan/export',
    name: 'LoanExport',
  },
  {
    path: '/loan/update',
    name: 'LoanUpdate',
  },
]

const router = new VueRouter({
  base: process.env.BASE_URL,
  routes
})

export default router
