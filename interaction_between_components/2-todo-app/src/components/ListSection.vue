<template>
    <TodoList/>
    <ResultList :itemCount="provideData.todoList.length"/>
</template>

<script>
import TodoList from "@/components/TodoList.vue"
import ResultList from "@/components/ResultList.vue"
export default {
    inject: ["provideData"],
    // props : {
    //     todoList: {
    //         type: Array,
    //         required: true
    //     },
    //     deleteItem: {
    //         type: Function,
    //         required: true
    //     }
    // },
    components: {
        TodoList,
        ResultList
    }
}
</script>