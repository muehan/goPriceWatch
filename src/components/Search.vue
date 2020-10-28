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
          <a class="waves-effect waves-light btn" v-on:click="search">Suche</a>
        </div>
      </form>
    </div>
    <template v-if="this.loaded">
      <Details :simplename="this.simplename" :name="this.name" :fullname="this.fullname" />
      <Price :id="this.productId" />
    </template>
  </div>
</template>

<script>
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
      productId: "",
      searchModel: "",
      loaded: false,
      name: "",
      simplename: "",
      fullname: ""
    };
  },
  methods: {
    search() {
      if (this.searchModel) {
        this.$http.get("/api/product/" + this.searchModel).then(
          response => {
            this.name = response.body.Name;
            this.simplename = response.body.SimpleName;
            this.fullname = response.body.Fullname;
            this.productId = response.body.Id;
            this.loaded = true;
            return response.body;
          },
          response => {
            this.loaded = false;
            console.error(response);
          }
        );
      }
    },
    submit(e) {
      this.search();
      e.preventDefault();
    }
  },
  mounted() {
    let productNumber = this.$route.params.productNumber;
    if (productNumber) {
      this.searchModel = productNumber;
      this.search();
    }
  }
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
