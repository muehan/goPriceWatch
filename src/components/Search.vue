<template>
  <div class="container">
    <h1>PriceWatch</h1>
    <div class="row">
      <form class="col s12">
        <div class="row">
          <div class="input-field col offset-s4 s4">
            <input placeholder="ProductId" id="productid" type="text" class="validate" v-model="search" />
            <label for="productid">ProductId</label>
          </div>
        </div>
        <div class="row">
          <a class="waves-effect waves-light btn" v-on:click="search">Suche</a>
        </div>
      </form>
    </div>
    <Details v-if="loaded" :simplename="this.simplename" :name="this.name" :fullname="this.fullname" />
  </div>
</template>

<script>
import Details from "./Details.vue";
import Vue from "vue";
import Resource from "vue-resource";
Vue.use(Resource);

export default {
  name: "Search",
  components: {
    Details
  },
  props: {},
  data() {
    return {
      search: '',
      loaded: true,
      name: "asdf",
      simplename: "",
      fullname: ""
    };
  },
  methods: {
    search() {
      console.log("search");
      this.$http.get("http://localhost:8080/api/product/8908274").then(
        response => {
          let data = JSON.parse(response.body);
          this.name = data.Name;
          this.simplename = data.SimpleName;
          this.fullname = data.Fullname;
          this.loaded = true;
          return response.body;
        },
        response => {
          console.error(response);
          this.name = "ertz";
        }
      );
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
