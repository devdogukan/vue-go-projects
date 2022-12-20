<template>
    <h1>Details</h1>
    <p>{{ $route.params.userID }}</p>
    <button @click="this.$router.push('/')">Back</button>
</template>