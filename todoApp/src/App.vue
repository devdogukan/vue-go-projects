<template>

  <form :action="sendTodo" @click.prevent="onsubmit">
    <div>
      <input v-model="title" type="text" class="form-control"/> 
      <button class="btn btn-primary" @click="sendTodo">Send</button>
    </div>
  </form>

  <v-container>
    <div class="input-group">
        <ul>
          <li class="list-group-item" v-for="(data, key) in todos" :key="key">
            {{ key+1 }}  ➡  {{data.title}}
            <button class="btn btn-primary" @click="updateTodo(data.id, true)">✔</button>
            <button class="btn btn-primary" @click="updateTodo(data.id, false)">✖</button>
            <button class="btn btn-primary" @click="deleteTodo(data.id, data.title, key)">🗑️</button>
          </li>
        </ul>
    </div>
  </v-container>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      title: "",
      socket: null,
      todos: [],
      count: 1
    }
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:9100/socket")
    this.socket.onmessage = (msg) => {
      this.addTodos(msg)
    }
  },
  methods : {
    sendTodo() {
      if (this.title != "") {
        let todo = {
        "procces": 0,
        "title" : this.title,
        "isDone" : false
        }
        let test = {
          id : this.count,
          title : this.title
        }
        this.socket.send(JSON.stringify(todo))
        this.todos.push(test)
        this.count++
        this.title = ""
      }
    },
    deleteTodo(id, title, key){
      let todo = {
        "procces": 1,
        "id" : id,
        "title" : title
      }
      this.socket.send(JSON.stringify(todo))
      this.todos.splice(key, 1)
    },
    addTodos(msg) {
      let todo = {
        id : this.count,
        title : msg.data
      }
      this.todos.push(todo)
      this.count++
    },
    updateTodo(id, isDone) {
      let todo = {
        "procces": 2,
        "id" : id,
        "isDone" : isDone
      }
      this.socket.send(JSON.stringify(todo))
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  align-items: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
