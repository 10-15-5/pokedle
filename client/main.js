import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'
import { loadFonts } from './plugins/webfontloader'
import PrimeVue from 'primevue/config';
import VirtualScroller from 'primevue/virtualscroller';
import Dialog from 'primevue/dialog';
import './style.css';
import 'primeicons/primeicons.css';


loadFonts();

createApp(App)
    .use(createPinia())
    .use(PrimeVue)
    .component('VirtualScroller', VirtualScroller)
    .component('Dialog', Dialog)
    .use(router)
    .mount('#app');
