import { createApp } from 'vue';
import App from './App.vue';
import plugins from './plugins';

import 'bootstrap/dist/css/bootstrap.min.css';
import './assets/styles/global.scss';

const app = createApp(App)
  .use(plugins.router)
  .use(plugins.store)
  .use(plugins.i18n)
  .use(plugins.ElementUI);

app.mount('#app');
