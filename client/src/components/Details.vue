<template>
  <div class="row">
    <div class="col s8 offset-s2">
      <div class="card blue-grey darken-1">
        <div class="card-content white-text">
          <span class="card-title">{{simplename}}</span>
          <p>{{name}}</p>
          <p>{{fullname}}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Details",
  props: {
    simplename: String,
    fullname: String,
    name: String
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
