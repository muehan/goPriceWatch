import Vue from 'vue';
import Vuex from 'vuex';
import search from './modules/search';
import productType from './modules/producttype';

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        search,
        productType,
    },
    strict: debug,
    // plugins: debug ? [createLogger()] : []
})