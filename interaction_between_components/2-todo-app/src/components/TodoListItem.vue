<template>
    <li>
        <span>{{ item.text }}</span>
        <button @click="deleteItem(item)" class="sm red">Delete</button>
      </li>
</template>

<script>
    export default {
        props : ["item"],
        inject: ["deleteItem"]
    }
</script>