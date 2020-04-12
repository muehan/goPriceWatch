<template>
  <div id="app">
    <nav class="teal">
      <div class="nav-wrapper">
        <router-link to="/" class="brand-logo center">Pricewatch</router-link>
        <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li>
             <router-link to="/">Search</router-link>
          </li>
          <li>
            <router-link to="/types">ProductTypes</router-link>
          </li>          
        </ul>
      </div>
    </nav>
    <div class="content">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  name: "App",
  components: {  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.content {
  margin-top: 60px;
}
</style>
