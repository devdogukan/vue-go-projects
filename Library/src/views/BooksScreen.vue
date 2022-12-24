<template>
    <div>
        <table style="width: 75%;">
            <tr>
                <th>Number</th>
                <th>Name</th>
                <th>Author</th>
                <th>Type</th>
                <th>Popularity</th>
                <th>Age Range</th>
            </tr>
            <tr v-for="(data, key) in books" :key="key">
                <td>{{ data.no }}</td>
                <td>{{ data.name }}</td>
                <td>{{ data.author }}</td>
                <td>{{ data.type }}</td>
                <td>{{ data.popularity }}</td>
                <td>{{ data.age_range }}</td>
            </tr>
        </table>
    </div>
</template>

<script>
export default {
    name: "BooksScreen",
    data() {
        return {
            socket: null,
            books: []
        }
    },
    mounted() {
        this.socket = new WebSocket("ws://localhost:9100/books")
        this.socket.onmessage = (msg) => {
            this.books = JSON.parse(msg.data)
        }
    }
}

</script>

<style>
    table,
    th,
    td {
        border: 1px solid black;
    }
    </style>