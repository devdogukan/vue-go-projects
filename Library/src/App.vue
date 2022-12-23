<template>
  <router-view></router-view>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      message : "Dre"
    }
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