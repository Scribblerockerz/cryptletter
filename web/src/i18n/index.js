import { createI18n } from 'vue-i18n';
import en from './en-messages.json';

export default createI18n({
    locale: 'en',
    legacy: false,
    messages: {
        en,
    },
});
