export default {
    loadProduct(productid) {
        return fetch(`/api/product/${productid}`);
    },

    loadPriceFor(productId, daysToLoad) {
        return fetch(`/api/price/${productId}/${daysToLoad}`);
    },

    loadProductsByType(typeid) {
        return fetch(`/api/productbytype/${typeid}`);
    },
}