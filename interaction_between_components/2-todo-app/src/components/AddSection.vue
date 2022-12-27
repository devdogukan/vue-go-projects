<template>
    <input @keydown.enter="addNewTodo($event.target.value)" type="text" id="todoText" placeholder="Add ToDo">
</template>

<script>
    export default {
        props : {
            addNewTodo: {
                type: Function,
                required: true
            }
        },
        methods: {
            // addNewTodo(event) {
            //     this.$emit("add-todo", event.target.value)
            //     event.target.value = ''
            // }
        }
    }
</script>