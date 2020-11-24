import productApi from '../../api/productApi';

const state = () => ({
    productId: "",
    product: {},
    prices: [],
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
    prices: (state) => {
        return state.prices
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
        commit('loadProduct', productId),
            productApi
            .loadProduct(productId)
            .then(response => response.json())
            .then(data => {
                commit('productLoadedSuccess', data);
            })
            .catch((error) => {
                commit('productLoadedFailed', error);
            });
    },
    loadPrices({ commit }, productId) {
        commit('loadPrices', productId),
            productApi
            .loadPriceFor(productId)
            .then(response => response.json())
            .then(data => {
                commit('priceLoadedSuccess', data);
            })
            .catch((error) => {
                commit('priceLoadedFailed', error);
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
    },

    loadPrices(state, productId) {
        state.productId = productId;
    },
    priceLoadedSuccess(state, prices) {
        state.product = prices;
    },
    priceLoadedFailed(state, error) {
        state.error = error;
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}