<template>
    <h1>Home</h1>
    <input type="text" @keydown.enter="gotoDetails"/>
</template>

<script>
export default {
    methods: {
        gotoDetails(e) {
            this.$router.push(`/details/${e.target.value}`)
        }
    }
}
</script>