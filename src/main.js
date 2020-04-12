import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router';
import Search from './components/Search.vue'
import ProductTypes from './components/ProductTypes.vue'

Vue.config.productionTip = false
Vue.use(VueRouter)

const routes = [
    { path: '/', component: Search },
    { path: '/types', component: ProductTypes }
]

const router = new VueRouter({
    routes // short for `routes: routes`
})

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')