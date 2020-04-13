<template>
  <div class="row">
    <div class="col offset-l2 m12 l8">
      <ul class="collection with-header">
        <li class="collection-header">
          <p>Folgende Produktegruppen können abgefragt werden:</p>
        </li>
        <li v-for="type in types" :key="type" class="collection-item">
          <div>
            {{type.Name}} | {{type.Typeid}}
            <router-link :to="'/productsbytype/' + type.Id" class="secondary-content">
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
  name: "ProductTypes",
  components: {},
  props: {},
  data() {
    return {
      types: []
    };
  },
  mounted() {
    this.loadTypes();
  },
  methods: {
    loadTypes() {
      this.$http.get("/api/producttype").then(
        response => {
          let data = JSON.parse(response.body);
          this.types = data;
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
