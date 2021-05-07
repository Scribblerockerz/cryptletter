<template>
    <div class="password-prompt" :class="{ 'password-prompt--open': isOpen }">
        <div
            :class="{ 'password-prompt__backdrop--active': isOpen }"
            class="password-prompt__backdrop"
            @click="handleCancel"
        />
        <form
            :class="{ 'password-prompt__box--active': isOpen }"
            class="password-prompt__box"
            @submit.prevent="handleSubmit"
        >
            <div class="password-prompt__head">
                <h3 class="password-prompt__headline">
                    {{ t("passwordProtection.headline") }}
                </h3>
                <p class="password-prompt__description">
                    {{ t("passwordProtection.description") }}
                </p>
            </div>
            <div class="password-prompt__body">
                <label for="postsage-stamp" class="password-prompt__label">
                    {{ t("passwordProtection.passwordLabel") }}
                </label>
                <input
                    ref="passwordInput"
                    v-model="password"
                    :disabled="!isOpen"
                    type="password"
                    class="password-prompt__input"
                />

                <div class="password-prompt__actions">
                    <Button
                        type="submit"
                        :disabled="!isPasswordValid || !isOpen"
                        primary
                        class="password-prompt__confirm-button"
                    >
                        {{ t("passwordProtection.confirmLabel") }}
                    </Button>
                    <Button
                        class="password-prompt__cancel"
                        :disabled="!isOpen"
                        @click="handleCancel"
                    >
                        {{ t("passwordProtection.cancelLabel") }}
                    </Button>
                </div>
            </div>
        </form>
    </div>
</template>

<script>
import { useI18n } from "vue-i18n";
import { ref } from "@vue/reactivity";
import { computed, onBeforeUnmount, onMounted, watch } from "@vue/runtime-core";
import Button from "../components/Button";

const isOpen = ref(false);
const password = ref("");
const resolveCallback = ref(null);
const rejectCallback = ref(null);

const resetPrompt = () => {
    document.body.classList.remove("js--password-prompt--open");
    isOpen.value = false;
    password.value = null;
    resolveCallback.value = null;
    rejectCallback.value = null;
};

export async function requestPasswordPrompt() {
    return new Promise((resolve, reject) => {
        document.body.classList.add("js--password-prompt--open");
        isOpen.value = true;
        resolveCallback.value = (result) => {
            resolve(result);
            resetPrompt();
        };
        rejectCallback.value = () => {
            reject();
            resetPrompt();
        };
    });
}

export default {
    name: "PasswordPrompt",
    components: {
        Button,
    },
    setup() {
        const { t } = useI18n();
        const isPasswordValid = computed(() => password.value !== "");
        const passwordInput = ref(null);

        const handleCancel = () => {
            rejectCallback.value && rejectCallback.value();
        };

        const handleSubmit = () => {
            if (!isPasswordValid.value) return;
            resolveCallback.value && resolveCallback.value(password.value);
        };

        const escapeAbortHandler = (e) => {
            if (e.key !== "Escape" || !isOpen.value) return;
            handleCancel();
        };

        onMounted(() => {
            document.addEventListener("keydown", escapeAbortHandler);
        });

        onBeforeUnmount(() => {
            document.removeEventListener("keydown", escapeAbortHandler);
        });

        // Trigger on open
        watch(isOpen, (value) => {
            if (!value) return;
            setTimeout(() => passwordInput.value.focus(), 100);
        });

        return {
            isOpen,
            password,
            isPasswordValid,
            passwordInput,

            handleSubmit,
            handleCancel,
            t,
        };
    },
};
</script>

<style>
.js--password-prompt--open {
    overflow: hidden;
}
</style>

<style scoped>
.password-prompt__box {
    position: fixed;
    top: 0;
    left: 50%;
    transform: translate(-50%, -100%);
    background-color: #fff;
    color: #222;
    opacity: 0;
    transition: opacity 0.3s ease-out, transform 0.3s ease-out;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
    border-bottom-left-radius: 4px;
    border-bottom-right-radius: 4px;
    width: 100%;
    max-width: 500px;
}

.password-prompt__box--active {
    transform: translate(-50%, 0);
    opacity: 1;
}

.password-prompt__head,
.password-prompt__body {
    padding: 20px;
}

.password-prompt__head > *:last-child {
    margin-bottom: 0;
}

.password-prompt__headline,
.password-prompt__description {
    margin-top: 0;
}

.password-prompt__body {
    background-color: #efefef;
}

.password-prompt__backdrop {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    pointer-events: none;
    background-color: rgba(0, 0, 0, 0.6);
    opacity: 0;
    transition: opacity 0.6s ease-out;
}

.password-prompt__backdrop--active {
    pointer-events: all;
    opacity: 1;
}

.password-prompt__label {
    display: block;
    font-weight: bold;
}

.password-prompt__input {
    width: 100%;
    padding: 10px;
    font-size: 18px;
    line-height: 1;
}

.password-prompt__actions {
    margin-top: 20px;
    display: flex;
    justify-content: space-between;
}
</style>
