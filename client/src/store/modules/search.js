import productApi from '../../api/productApi';

const state = () => ({
    productId: "",
    product: {},
    error: "",
    loaded: false,
})

// getters
const getters = {
    productId: (state) => {
        return state.productId
    },
    product: (state) => {
        return state.product
    },
    // error: (state, getters, rootState) => {
    error: (state) => {
        return state.error
    },
    loaded: (state) => {
        return state.loaded
    },
}

const actions = {
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
}

// mutations
const mutations = {
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
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}