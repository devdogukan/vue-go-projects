<template>
    <!-- <ul>
      <li v-for="todoItem in todoList" :key="todoItem.id">
        <span>{{ todoItem.text }}</span>
        <button @click="deleteItem(todoItem)" class="sm red">Delete</button>
      </li>
    </ul> -->

    <ul>
      <TodoListItem @delete-todo-item="$emit('delete-todo-item', $event)" v-for="todoItem in provideData.todoList" :key="todoItem.id" :item="todoItem"/>
    </ul>
</template>

<script>
import TodoListItem from "@/components/TodoListItem.vue"
    export default {
        inject: ["provideData"],
        components: {
            TodoListItem
        },
        //props : ["todoList"]
    }
</script>