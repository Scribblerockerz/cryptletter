<template>
    <div :class="{ 'page--unavailable': isUnavailable }" class="page">
        <div class="page__header">
            <h1 class="page__title">
                <a class="logo" href="/">Cryptletter</a>
            </h1>
            <div class="page__slogan">
                {{ t("slogan") }}
            </div>
        </div>
        <div class="page__container">
            <Toaster />
            <slot />
        </div>
    </div>
</template>

<script>
import { useI18n } from "vue-i18n";
import Toaster from "./Toaster";

export default {
    name: "Page",
    components: {
        Toaster,
    },
    props: {
        isUnavailable: {
            type: Boolean,
            default: false,
        },
    },
    setup() {
        const { t } = useI18n();
        return { t };
    },
};
</script>

<style scoped>
.page {
    margin-bottom: 70px;
}

.page__header {
    text-align: center;
    line-height: 1.2em;
    margin: 70px auto;
}

.page__slogan {
    white-space: pre;
}

.page__container {
    margin: auto;
    max-width: 840px;
    padding: 0 20px;
}

.page--unavailable .page__header {
    position: absolute;
    bottom: 70px;
    margin: 0 auto;
    width: 100%;
}

.page--unavailable .logo {
    opacity: 0.1;
    transition: opacity 0.2s ease-in-out;
    background-size: 140px;
}

.page--unavailable .logo:hover,
.page--unavailable .logo:focus {
    opacity: 1;
}

.page--unavailable .page__slogan {
    display: none;
}

.page--unavailable .page__container {
    max-width: 540px;
    width: 100%;
    position: absolute;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
}
</style>
