import { createApp } from 'vue'
import App from './App.vue'
import appHeader from "@/components/appHeader.vue"

const app = createApp(App)

app.component("app-header", appHeader)

app.mount('#app')
