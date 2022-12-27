<template>
  <div>
    <app-header/>
    <CounterItem/>
    <CounterItem/>
  </div>
  
</template>

<script>
  import CounterItem from "@/components/CounterApp.vue";

  export default {
    components: {
      CounterItem
    }
  }
</script>