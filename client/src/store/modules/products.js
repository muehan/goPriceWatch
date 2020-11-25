import productApi from '../../api/productApi';

const state = () => ({
    products: [],
    error: "",
    loaded: false,
})

// getters
const getters = {
    products: (state) => state.products,
    productsFiltered: (state) => (term) => state.products.filter(x => x.Fullname.toLowerCase().indexOf(term.toLowerCase()) >= 0),
    error: (state) => state.error,
    loaded: (state) => state.loaded,
}

const actions = {
    loadProducts({ commit }, typeId) {
        commit('loadProducts'),
            productApi
            .loadProductsByType(typeId)
            .then(response => response.json())
            .then(data => {
                commit('loadProductsSuccess', data);
            })
            .catch((error) => {
                commit('loadProductsFailed', error);
            });
    }
}

// mutations
const mutations = {
    loadProducts(state, ) {
        state.loaded = false;
    },
    loadProductsSuccess(state, products) {
        state.products = products;
        state.loaded = true;
    },
    loadProductsFailed(state, error) {
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