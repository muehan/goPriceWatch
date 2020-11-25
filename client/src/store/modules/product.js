import productApi from '../../api/productApi';
import moment from "moment";

const state = () => ({
    productId: "",
    product: {},
    prices: [],
    error: "",
    loaded: false,
    daysToLoad: 30,
})

// getters
const getters = {
    productId: (state) => state.productId,
    product: (state) => state.product,
    prices: (state) => state.prices,
    // pricesValues: (state) => state.prices.map(x => x.Price),
    labels: (state) => state.prices.map(x => moment(String(x.Date)).format("YYYY-MM-DD")),
    daysToLoad: (state) => state.daysToLoad,
    // error: (state, getters, rootState) => {
    error: (state) => state.error,
    loaded: (state) => state.loaded,
    minPrice: (state) => Math.min(...state.prices.map(x => x.Price)),
    maxPrice: (state) => Math.max(...state.prices.map(x => x.Price)),
    datacollection: (state) => {
        return {
            labels: state.prices.map(x =>
                moment(String(x.Date)).format("YYYY-MM-DD")
            ),
            datasets: [{
                label: "Price",
                backgroundColor: "rgba(153, 159, 255, 0.2)",
                data: state.prices.map(x => x.Price),
            }]
        };
    }
}

const actions = {
    setDaysToLoad({ commit, dispatch }, daysToLoad) {
        commit('setDaysToLoad', daysToLoad);
        dispatch('loadPrices');
    },
    loadProduct({ commit, dispatch }, productId) {
        commit('loadProduct', productId),
            productApi
            .loadProduct(productId)
            .then(response => response.json())
            .then(data => {
                commit('productLoadedSuccess', data);
                dispatch('loadPrices');
            })
            .catch((error) => {
                commit('productLoadedFailed', error);
            });
    },
    loadPrices({ commit, getters }) {
        commit('loadPrices'),
            productApi
            .loadPriceFor(getters.product.Id, getters.daysToLoad)
            .then(response => response.json())
            .then(data => {
                commit('priceLoadedSuccess', data);
            })
            .catch((error) => {
                commit('priceLoadedFailed', error);
            });
    },
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
        state.prices = prices;
    },
    priceLoadedFailed(state, error) {
        state.error = error;
    },
    setDaysToLoad(state, days) {
        state.daysToLoad = days;
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}