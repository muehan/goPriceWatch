<template>
  <div>
    <div class="row">
      <div class="input-field col offset-l2 s12 m12 l8">
        <input id="search" type="text" class="validate" v-model="filter" />
        <label for="search">Search</label>
      </div>
    </div>
    <div class="row">
      <div class="col offset-l2 m12 l8" v-if="loaded">
        <ul class="collection with-header">
          <li class="collection-header">
            <p>Produkte</p>
          </li>
          <li
            v-for="product in productsFiltered(filter)"
            :key="product"
            class="collection-item"
          >
            <div>
              <div style="width: 85%; display: inline-block;">
                {{ product.Fullname }}
              </div>
              <div style="float: right;">
                <router-link
                  :to="'/search/' + product.ProductidAsString"
                  class="secondary-content link"
                >
                  <i class="material-icons">send</i>
                </router-link>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "ProductsByType",
  components: {},
  props: {},
  data() {
    return {
      filter: "",
    };
  },
  mounted() {
    this.loadProducts(this.$route.params.id);
  },
  methods: {
    ...mapActions({
      loadProducts: "products/loadProducts",
    }),
  },
  computed: {
    ...mapGetters({
      products: "products/products",
      productsFiltered: "products/productsFiltered",
      loaded: "products/loaded",
    }),
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.link {
  height: 24px;
}
</style>
