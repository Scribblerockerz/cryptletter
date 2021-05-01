import { createI18n } from 'vue-i18n';
import en from './en-messages.json';

var activeLocale = 'en';
const messages = {
    en,
};

// Extendable locale configuration
if (window.cryptletterExtended) {
    // change active locale, at initialization
    activeLocale = window.cryptletterExtended.activeLocale || activeLocale;

    const extendedLocales = window.cryptletterExtended.locales || {};

    Object.keys(extendedLocales).forEach((localeKey) => {
        messages[localeKey] = messages[localeKey] || {};
        messages[localeKey] = {
            ...messages[localeKey],
            ...extendedLocales[localeKey],
        };
    });
}

export default createI18n({
    locale: activeLocale,
    legacy: false,
    messages: messages,
});
