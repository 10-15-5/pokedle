import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import { createPinia } from 'pinia'
import { loadFonts } from './plugins/webfontloader'
import PrimeVue from 'primevue/config';
import VirtualScroller from 'primevue/virtualscroller';
import './style.css';

loadFonts();

createApp(App)
    .use(createPinia())
    .use(PrimeVue)
    .component('VirtualScroller', VirtualScroller)
    .use(router)
    .use(vuetify)
    .mount('#app');
