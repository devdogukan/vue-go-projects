<template>
  <div class="row">
    <div class="col-sm-1"></div>
    <div class="col-sm-4">
      <canvas id="barChart" width="400" height="400"></canvas>
    </div>
    <div class="col-sm-1"></div>
    <div class="col-sm-4">
      <canvas id="pieChart" width="400" height="400"></canvas>
    </div>
  </div>
</template>

<script>
import Chart from "chart.js/auto";

export default {
  name: 'HelloWorld',
  props: {
    msg: String,
  },
  data() {
    return {
      socket: null,
      book_names: [],
      popularity: [],
      book_types: [],
      book_types_count: []
    }
  },
  beforeMount() {
    this.socket = new WebSocket("ws://localhost:9100/bar")
    this.socket.onmessage = (msg) => {
      var message = JSON.parse(msg.data)
      for (let i = 0; i < message.length; i++) {
        this.book_names.push(message[i].name)
        this.popularity.push(message[i].popularity)
        if (!this.book_types.includes(message[i].type)) {
          this.book_types.push(message[i].type)
          this.book_types_count.push(1)
        } else {
          var index = this.book_types.indexOf(message[i].type)
          this.book_types_count[index]++
        }
      }
    }
  },

  mounted() {

    const barChart = document.getElementById("barChart")
    const pieChart = document.getElementById("pieChart")

    setTimeout(() => {

      new Chart(barChart, {
        type: 'bar',
        data: {
          labels: this.book_names,
          datasets: [{
            label: '# Book Popularity',
            data: this.popularity,
            backgroundColor: [
              'rgba(255, 99, 132, 0.2)',
              'rgba(54, 162, 235, 0.2)',
              'rgba(255, 206, 86, 0.2)',
              'rgba(75, 192, 192, 0.2)',
              'rgba(153, 102, 255, 0.2)',
              'rgba(255, 159, 64, 0.2)'
            ],
            borderColor: [
              'rgba(255,99,132,1)',
              'rgba(54, 162, 235, 1)',
              'rgba(255, 206, 86, 1)',
              'rgba(75, 192, 192, 1)',
              'rgba(153, 102, 255, 1)',
              'rgba(255, 159, 64, 1)'
            ],
            borderWidth: 1
          }]
        },
        options: {
          scales: {
            y: {
              beginAtZero: true
            }
          }
        }
      });

      new Chart(pieChart, {
        type: 'pie',
        data: {
          labels: this.book_types,
          datasets: [{
            label: "Book Type Dataset",
            data: this.book_types_count,
            backgroundColor: [
              'rgb(255, 99, 132)',
              'rgb(54, 162, 235)',
              'rgb(255, 205, 86)'
            ],
            hoverOffset: 4
          }]
        }
      });

    }, 100);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
