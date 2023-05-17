import { createApp } from 'vue'
import App from './App.vue'
import './index.css'
import 'jquery';
import 'bootstrap/dist/css/bootstrap.min.css';
import * as bootstrap from 'bootstrap';
window.bootstrap = bootstrap;

import VueTyped from 'vue3-typed-js';

const app = createApp(App);
app.use(VueTyped);
app.mount('#app');

