export default {
    loadProduct(productid) {
        return fetch(`/api/product/${productid}`);
    },

    loadPriceFor(id, daysToLoad) {
        return fetch(`/api/price/${id}/${daysToLoad}`);
    },

    loadProductsByType(typeid) {
        return fetch(`/api/productbytype/${typeid}`);
    },
}