<template>
  <div>
    <div class="row">
      <div class="input-field col offset-l2 s12 m12 l8">
        <input id="search" type="text" class="validate" v-model="filter" />
        <label for="search">Search</label>
      </div>
    </div>
    <div class="row">
      <div class="col offset-l2 m12 l8">
        <ul class="collection with-header">
          <li class="collection-header">
            <p>Produkte</p>
          </li>
          <li v-for="product in filteredProducts()" :key="product" class="collection-item">
            <div>
              <div style="width: 85%; display: inline-block;">{{product.Fullname}}</div>
              <div style="float: right;">
                <router-link
                  :to="'/search/' + product.ProductidAsString"
                  class="secondary-content link">
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
import Vue from "vue";
import Resource from "vue-resource";
Vue.use(Resource);

export default {
  name: "ProductsByType",
  components: {},
  props: {},
  data() {
    return {
      products: [],
      filter: ""
    };
  },
  mounted() {
    this.loadTypes();
  },
  methods: {
    loadTypes() {
      this.$http.get("/api/productbytype/" + this.$route.params.id).then(
        response => {
          let data = JSON.parse(response.body);
          this.products = data;
          return response.body;
        },
        response => {
          console.error(response);
        }
      );
    },
    filteredProducts() {
      return this.products.filter(x => {
          return x
            .Name
            .toLowerCase()
            .indexOf(
              this.filter.toLowerCase()) >= 0
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.link {
  height: 24px;
}
</style>
