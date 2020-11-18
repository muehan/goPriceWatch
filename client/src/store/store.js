import Vue from 'vue'
import Vuex from 'vuex'
import productApi from '../api/productApi';

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        productId: "",
        product: {},
        error: "",
        loaded: false,
    },
    mutations: {
        loadProduct(state, productId) {
            state.productId = productId;
            state.loaded = false;
        },
        productLoadedSuccess(state, product) {
            state.product = product;
            state.loaded = true;
        },
        productLoadedFailed(state, error) {
            state.error = error;
            state.loaded = false;
        }
    },
    actions: {
        loadProduct({ commit }, productId) {
            console.log(`action loadProduct with ${productId}`);
            commit('loadProduct', productId),
                productApi
                .loadProduct(productId)
                .then(response => response.json())
                .then(data => {
                    console.log(`loadProduct from api`);
                    console.log(data);
                    commit('productLoadedSuccess', data);
                })
                .catch((error) => {
                    commit('productLoadedFailed', error);
                });
        }
    },
    getters: {
        productId: state => state.productId,
        product: state => state.product,
        error: state => state.error,
        loaded: state => state.loaded,
    }
})