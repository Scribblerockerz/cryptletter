import { createApp } from 'vue';
import i18n from './i18n';
import App from './App.vue';

if (process.env.NODE_ENV === 'development') {
    window.cryptletterOptions = {
        supportsAttachments: true,
    };
}

createApp(App).use(i18n).mount('#app');

// TODO: https://stackoverflow.com/questions/64540998/vue-js-3-error-while-using-vuei18n-plugin-in-vue-cannot-set-property-vm-o/64631506#64631506
