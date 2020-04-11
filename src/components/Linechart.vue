<script>
import { Line } from "vue-chartjs";

export default {
  name: "Linechart",
  extends: Line,
  props: ["chartdata", "options"],
  mounted() {
    this.renderChart(this.chartdata, this.options);
  }
};
</script>

<style>
</style>