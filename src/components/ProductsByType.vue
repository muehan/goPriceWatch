<template>
  <div class="row">
    <div class="col offset-l2 m12 l8">
      <ul class="collection with-header">
        <li class="collection-header">
          <p>Produkte</p>
        </li>
        <li v-for="product in products" :key="product" class="collection-item">
          <div>
            {{product.Name}}
            <router-link :to="'/search/' + product.ProductidAsString" class="secondary-content">
              <i class="material-icons">send</i>
            </router-link>
          </div>
        </li>
      </ul>
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
      products: []
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
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
