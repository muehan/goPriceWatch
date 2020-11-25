<template>
  <div class="container">
    <div style="margin-bottom: 40px">
      <p>Show days</p>
      <a v-on:click="changeDaysToLoad(30)" class="waves-effect waves-light btn">30</a>
      <a v-on:click="changeDaysToLoad(60)" class="waves-effect waves-light btn">60</a>
      <a v-on:click="changeDaysToLoad(90)" class="waves-effect waves-light btn">90</a>
      <a v-on:click="changeDaysToLoad(180)" class="waves-effect waves-light btn">180</a>
      <a v-on:click="changeDaysToLoad(0)" class="waves-effect waves-light btn">max</a>
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
import { mapGetters, mapActions } from "vuex";
import moment from "moment";
import LineChart from "./LineChart.js";

export default {
  name: "Price",
  props: { },
  components: {
    LineChart
  },
  data() {
    return {
      options: {
        responsive: true,
        maintainAspectRatio: false
      }
    };
  },
  computed: {
    ...mapGetters({
      prices: 'product/prices',
      minPrice: 'product/minPrice',
      maxPrice: 'product/maxPrice',
      pricesValues: 'product/pricesValues',
      datacollection: 'product/datacollection',
    }),
  },
  mounted() { },
  methods: {
    ...mapActions({
      changeDaysToLoad: 'product/setDaysToLoad',
    }),
    format_date(value) {
      if (value) {
        return moment(String(value)).format("YYYY-MM-DD");
      }
    },
  },
  created: function() {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
