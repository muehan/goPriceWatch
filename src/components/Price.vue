<template>
  <div class="container">
    <ul>
      <li v-for="price in prices" :key="price.Price">{{format_date(price.Date)}} | {{price.Price}}.-</li>
    </ul>
  </div>
</template>

<script>
import moment from "moment";
import Vue from "vue";
import Resource from "vue-resource";
Vue.use(Resource);

export default {
  name: "Price",
  props: {
    id: String
  },
  data() {
    return {
      productId: this.id,
      prices: []
    };
  },
  methods: {
    format_date(value) {
      if (value) {
        return moment(String(value)).format("YYYY-MM-DD");
      }
    }
  },
  computed: {},
  watch: {
    id: {
      immediate: true,
      handler(newValue, oldValue) {
        console.log(oldValue + " - " + newValue);
        this.$http
          .get("http://localhost:8080/api/price/" + newValue)
          .then(
            function(response) {
              let data = JSON.parse(response.body);
              console.log(data);
              this.prices = data;
            },
            function(response) {
              console.error(response);
            }
          );
      }
    }
  },
  created: function() {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
