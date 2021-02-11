import { createApp } from 'vue';
import VueI18n from 'vue-i18n';
import App from './App.vue';

const app = createApp(App);
app.use(VueI18n);

app.mount('#app');
