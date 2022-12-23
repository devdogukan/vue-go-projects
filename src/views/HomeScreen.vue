<template>
    <div class="chartjs">
        <div class="col-lg-4 grid-margin stretch-card">
          <p>{{ datacollection.labels }}</p>
          <p>{{ datacollection.datasets[0].data }}</p>
            <Bar
                id="my-chart-id"
                :options="options"
                :data="datacollection"
            />
        </div>
    </div>
  </template>
  
  <script>
  import { Bar } from 'vue-chartjs';
  import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js';
  
  ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)
  
  export default {
    name: 'BarChart',
    extends: Bar,
    components: { Bar },
    data() {
      return {
        socket: null,
        datacollection: {
        labels: [],
        datasets: [{
          label: '# of Votes',
          data: [],
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
          yAxes: [{
            ticks: {
              beginAtZero: true
            },
            gridLines: {
              display: false
            }
          }],
          xAxes: [{
            ticks: {
              beginAtZero: true
            },
            gridLines: {
              display: false
            }
          }]
        },
        legend: {
          display: false
        },
        elements: {
          point: {
            radius: 0
          }
        }
      }
    }
  },
  // mounted() {
  //   this.socket = new WebSocket("ws://localhost:9100/socket")
  //   this.socket.onmessage = (msg) => {
  //     var message = JSON.parse(msg.data)
  //     this.datacollection.labels.push(message.name)
  //     this.datacollection.datasets[0].data.push(message.popularity)
  //     console.log(this.datacollection.labels)
  //   }
  // },
  mounted () {
    this.renderChart(this.datacollection, this.options)
    this.socket = new WebSocket("ws://localhost:9100/socket")
    this.socket.onmessage = (msg) => {
      var message = JSON.parse(msg.data)
      this.datacollection.labels.push(message.name)
      this.datacollection.datasets[0].data.push(message.popularity)
      console.log(this.datacollection.labels)
    }
  }
}
</script>