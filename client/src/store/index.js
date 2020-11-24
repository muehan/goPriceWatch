import Vue from 'vue';
import Vuex from 'vuex';
import product from './modules/product';
import products from './modules/products';
import productType from './modules/producttype';

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        product,
        products,
        productType,
    },
    strict: debug,
    // plugins: debug ? [createLogger()] : []
})