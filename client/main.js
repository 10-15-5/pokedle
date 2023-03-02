import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import { createPinia } from 'pinia'
import { loadFonts } from './plugins/webfontloader'
import './style.css';
import PrimeVue from 'primevue/config';

loadFonts();

createApp(App)
    .use(createPinia())
    .use(PrimeVue)
    //.component('Card', Card)
    .use(router)
    .use(vuetify)
    .mount('#app');
