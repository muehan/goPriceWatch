import productTypeApi from '../../api/productTypeApi';

const state = () => ({
    producttypes: [],
    error: "",
    loaded: false,
})

// getters
const getters = {
    producttypes: (state) => {
        return state.producttypes
    },
    error: (state) => {
        return state.error
    },
    loaded: (state) => {
        return state.loaded
    },
}

const actions = {
    loadProducttypes({ commit }) {
        commit('loadProductTypes'),
            productTypeApi
            .load()
            .then(response => response.json())
            .then(data => {
                commit('productTypesLoadedSuccess', data);
            })
            .catch((error) => {
                commit('productTypesLoadedFailed', error);
            });
    }
}

// mutations
const mutations = {
    loadProductTypes(state, ) {
        state.loaded = false;
    },
    productTypesLoadedSuccess(state, producttypes) {
        state.producttypes = producttypes;
        state.loaded = true;
    },
    productTypesLoadedFailed(state, error) {
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