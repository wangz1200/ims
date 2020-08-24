import Vue from 'vue/dist/vue.js';
const { default: VueRouter } = require("vue-router")

Vue.use(VueRouter);

const Foo = { template: '<div>foo</div>' }

const routes = [
    { path: '/loan/query', component: Foo },
]

const router = new VueRouter({
    routes,
});

export default router;