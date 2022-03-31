import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'
import ElementPlus from 'element-plus'

import './assets/styles/global.scss'

createApp(App).use(store).use(router).use(ElementPlus).mount('#app')
