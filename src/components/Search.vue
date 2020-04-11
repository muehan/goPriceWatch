<template>
  <div>
    <h1>PriceWatch</h1>
    <div class="row">
      <form class="col s12" @submit="submit">
        <div class="row">
          <div class="input-field col offset-s4 s4">
            <input
              placeholder="ProductId"
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
            let data = JSON.parse(response.body);
            this.name = data.Name;
            this.simplename = data.SimpleName;
            this.fullname = data.Fullname;
            this.productId = data.Id;
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
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
</style>
