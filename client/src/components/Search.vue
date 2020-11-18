<template>
  <div>
    <div class="row">
      <p>Hier Produktnummer aus Digitec Shop eintragen:</p>
      <p class="small">
        Folgende Produkttypen werden gespeichert:
        <router-link to="/types">typen</router-link>
      </p>
      <form class="col s12" @submit="submit">
        <div class="row">
          <div class="input-field col offset-s4 s4">
            <input
              id="productid"
              type="text"
              class="validate"
              v-model="searchModel"
            />
            <label for="productid">ProductId</label>
          </div>
        </div>
        <div class="row">
          <a class="waves-effect waves-light btn" v-on:click="loadProduct(searchModel)">Suche</a>
        </div>
      </form>
    </div>
    <template v-if="this.loaded">
      <Details/>
      <Price :id="this.product.Id" />
    </template>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Details from "./Details.vue";
import Price from "./Price.vue";
import Vue from "vue";
import Resource from "vue-resource";
Vue.use(Resource);

export default {
  name: "Search",
  components: {
    Details,
    Price
  },
  props: {},
  data() {
    return {
      searchModel: "",
    };
  },
  methods: {
    // search() {
    //   if (this.searchModel) {
    //     this.$store.dispatch('loadProduct', this.searchModel);
    //   }
    // },
    submit(e) {
      this.$store.dispatch('loadProduct', this.searchModel);
      e.preventDefault();
    },
    ...mapActions([
      'loadProduct',
    ])
  },
  mounted() {
    let productNumber = this.$route.params.productNumber;
    if (productNumber) {
      this.searchModel = productNumber;
      this.$store.dispatch('loadProduct', this.searchModel);
    }
  },
  computed: mapState({
    product: state => state.product,
    loaded: state => state.loaded,
  }),
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
.small {
  font-size: 10px;
}
</style>
