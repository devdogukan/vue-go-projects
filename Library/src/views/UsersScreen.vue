<template>

    <h1>Users Page</h1>

    <div class="col-sm-8">
        <pre> ▶ Access permission 1: Permission to access books <b>under</b> the age of 18</pre>
        <pre> ▶ Access permission 2: Permission to access books <b>over</b> the age of 18</pre>
    </div>

    <div class="row">
        <div class="col-sm-8">
            <table style="width: 100%;">
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Surname</th>
                    <th>Age</th>
                    <th>Access permission</th>
                </tr>
                <tr v-for="(user, key) in users" :key="key">
                    <td>{{ user.id }}</td>
                    <td>{{ user.name }}</td>
                    <td>{{ user.surname }}</td>
                    <td>{{ user.age }}</td>
                    <td v-if="user.access_permission == true">
                        <input type="checkbox" :value="user.access_permission" checked @click="changeAcces(user)">
                    </td>
                    <td v-if="user.access_permission == false">
                        <input type="checkbox" :value="user.access_permission" @click="changeAcces(user)">
                    </td>
                </tr>
            </table>
            <button class="btn btn-outline-success" type="button" @click="save()">Save</button>
        </div>
    </div>
</template>

<script>
export default {
    name: "BooksScreen",
    data() {
        return {
            socket: null,
            users: [],
            changeAccesList: []
        }
    },
    mounted() {
        this.socket = new WebSocket("ws://localhost:9100/users")
        this.socket.onmessage = (msg) => {
            this.users = JSON.parse(msg.data)
        }
    },
    methods : {
        changeAcces(user) {
            user.access_permission = !user.access_permission

            if (!this.changeAccesList.includes(user)) {
                this.changeAccesList.push(user)
            } else {
                var index = this.changeAccesList.indexOf(user)
                this.changeAccesList[index] = user
            }
        },
        save() {
            if(this.changeAccesList.length > 0) {
                this.socket.send(JSON.stringify(this.changeAccesList))
                this.changeAccesList.length = 0
            }
        }
    }
}

</script>

<style>
    table,
    th,
    td {
        border: 1px solid black;
        text-align: center;
    }
</style>