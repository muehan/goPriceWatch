export default {
    loadProduct(productid) {
        return fetch("/api/product/" + productid);
    }
}