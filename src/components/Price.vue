<template>
  <div class="container">
    <line-chart :chart-data="datacollection" :options="options"></line-chart>
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

import LineChart from "./LineChart.js";

export default {
  name: "Price",
  props: {
    id: String
  },
  components: {
    LineChart
  },
  data() {
    return {
      productId: this.id,
      prices: [],
      dates: [],
      datacollection: null,
      options: {
          responsive: true,
          maintainAspectRatio: false,
      },
    };
  },
  computed: {
      
  },
  mounted() {
    // this.fillData();
  },
  methods: {
    format_date(value) {
      if (value) {
        return moment(String(value)).format("YYYY-MM-DD");
      }
    },
    fillData() {
      this.datacollection = {
        labels: this.prices.map(x => moment(String(x.Date)).format("YYYY-MM-DD")),
        datasets: [
          {
            label: "Price",
            backgroundColor: "rgba(153, 159, 255, 0.2)",
            data: this.prices.map(x => x.Price),
          },
        ]
      };
    },
    getRandomInt() {
      return Math.floor(Math.random() * (50 - 5 + 1)) + 5;
    }
  },
  watch: {
    id: {
      immediate: true,
      handler(newValue, oldValue) {
        console.log(oldValue + " - " + newValue);
        this.$http.get("/api/price/" + newValue).then(
          function(response) {
            let data = JSON.parse(response.body);
            this.prices = data;
            this.fillData();
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
