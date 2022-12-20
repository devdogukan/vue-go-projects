<template>
  <div class="container">
    <div class="mb-2">
      <router-link class="nav-link" active-class="active" to="/">
      Home
      </router-link>
      <router-link class="nav-link" active-class="active" to="/about">
        About
      </router-link>
      <router-link class="nav-link" active-class="active" to="/details/123">
        Detail
      </router-link>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
  export default {
    created(){
      console.log(this.$route)
    }
  }
</script>

<style>
  .nav-link {
    padding: 5px 10px;
    text-decoration: none;
    border: 1px solid #fa6558;
    color: #fa6558;
    margin-right: 5px;
  }
  .active {
    background-color: #fa6558;
    color: #fff;
  }
</style>