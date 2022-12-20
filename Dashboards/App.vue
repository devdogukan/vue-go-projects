<template>
  <div class="app">
    <h1>Hello, World!</h1>
    <HomeScreenVue></HomeScreenVue>
  </div>
</template>

<script>
import HomeScreenVue from './views/HomeScreen.vue';

export default {
  name: 'App',
  components: {
    HomeScreenVue
  }
}
</script>

<style lang="scss">

</style>
