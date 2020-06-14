<template>
  <div class="container">
    <div style="margin-bottom: 40px">
      <p>Show days</p>
      <a v-on:click="changeDaysToLoad(30)" class="waves-effect waves-light btn">30</a>
      <a v-on:click="changeDaysToLoad(60)" class="waves-effect waves-light btn">60</a>
      <a v-on:click="changeDaysToLoad(90)" class="waves-effect waves-light btn">90</a>
    </div>

    <div class="row">
      <line-chart :chart-data="datacollection" :options="options"></line-chart>
    </div>

    <div class="row">
      <div class="col s6 offset-s3">
        <h5>min: {{ minPrice }}.-</h5>
        <h5>max: {{ maxPrice }}.-</h5>
      </div>
    </div>

    <ul class="collection">
      <li
        class="collection-item"
        style="text-align: left"
        v-for="price in prices"
        :key="price.Price"
      >
        <span>{{format_date(price.Date)}}</span>
        <span v-if="price.InsteadOfPrice" style="float: right; color: red;">
          <span>{{price.Price}}.- instead of {{price.InsteadOfPrice}}.-</span>
        </span>
        <span v-else style="float: right;">
          <span>{{price.Price}}.-</span>
        </span>
      </li>
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
      daysToLoad: 30,
      productId: this.id,
      prices: [],
      dates: [],
      minPrice: 0,
      maxPrice: 0,
      datacollection: null,
      options: {
        responsive: true,
        maintainAspectRatio: false
      }
    };
  },
  computed: {},
  mounted() {
    // this.fillData();
  },
  methods: {
    changeDaysToLoad(days) {
      this.daysToLoad = days;
      this.loadData(this.productId);
    },
    format_date(value) {
      if (value) {
        return moment(String(value)).format("YYYY-MM-DD");
      }
    },
    fillData() {
      this.datacollection = {
        labels: this.prices.map(x =>
          moment(String(x.Date)).format("YYYY-MM-DD")
        ),
        datasets: [
          {
            label: "Price",
            backgroundColor: "rgba(153, 159, 255, 0.2)",
            data: this.prices.map(x => x.Price)
          }
        ]
      };
    },
    getRandomInt() {
      return Math.floor(Math.random() * (50 - 5 + 1)) + 5;
    },
    loadData(productId) {
      this.$http.get("/api/price/" + productId + "/" + this.daysToLoad).then(
        function(response) {
          let data = JSON.parse(response.body);
          this.prices = data;
          this.fillData();
          this.minPrice = Math.min(...this.prices.map(x => x.Price));
          this.maxPrice = Math.max(...this.prices.map(x => x.Price));
        },
        function(response) {
          console.error(response);
        }
      );
    }
  },
  watch: {
    id: {
      immediate: true,
      handler(newValue, oldValue) {
        console.log(oldValue + " - " + newValue);
        this.loadData(newValue);
      }
    }
  },
  created: function() {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
