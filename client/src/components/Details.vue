<template>
  <div class="row">
    <div class="col s8 offset-s2">
      <div class="card blue-grey darken-1">
        <div class="card-content white-text">
          <span class="card-title">{{product.SimpleName}}</span>
          <p>{{product.Name}}</p>
          <p>{{product.Fullname}}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  name: "Details",
  props: { },
  computed: mapGetters({
    product: 'product/product',
  }),
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
