import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import store from './store'
import Search from './components/Search.vue'
import ProductTypes from './components/ProductTypes.vue'
import ProductsByType from './components/ProductsByType.vue'

Vue.config.devtools = true;
Vue.config.productionTip = false;
Vue.use(VueRouter)

const routes = [
    { path: '/', component: Search },
    { path: '/search/:productNumber?', component: Search },
    { path: '/types', component: ProductTypes },
    { path: '/productsbytype/:id', component: ProductsByType },
]

const router = new VueRouter({
    routes // short for `routes: routes`
})

new Vue({
    router,
    store,
    render: h => h(App),
}).$mount('#app')