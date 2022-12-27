<template>
    <!-- <small class="mt-2">{{ todoList.length }}</small> -->
    <small class="mt-2">{{ itemCount }}</small>
</template>

<script>
    export default {
        props : ["itemCount"]
    }
</script>