<template>
    <Page :is-unavailable="isMissing" v-if="!isPending && isMissing">
        <h2>{{ t("missingMessageHeadline") }}</h2>
        <p>{{ t("missingMessageText") }}</p>
    </Page>
    <Page v-if="!isPending && !isMissing">
        <Letter striped spaced foot-separator>
            <template v-slot:header>
                <Button v-if="!showRaw" nano @click="toggleMode">
                    view raw
                </Button>
                <Button v-else nano @click="toggleMode">view text</Button>
            </template>
            <TextareaField
                v-if="showRaw"
                :modelValue="message"
                readonly
                raw
                focus-selection
                :rows="10"
                :placeholder="
                    isPending ? t('decryptingMessage') : t('messageEmpty')
                "
            />
            <pre
                v-else
            ><code id="message">{{ isPending ? t('decryptingMessage') : message ? message : t('messageEmpty') }}</code></pre>

            <template v-slot:footer>
                <div class="u--danger">
                    {{
                        t("estimatedLifetimeHint", {
                            duration: readableDuration,
                        })
                    }}
                </div>
            </template>
        </Letter>
        <div class="u--center">
            <Button
                :disabled="isDestroyPending"
                primary
                @click="destroyMessage"
                >{{ t("destroyMessageLabel") }}</Button
            >
        </div>
    </Page>
</template>

<script>
import Page from "../components/Page";
import Letter from "../components/Letter";
import TextareaField from "../components/TextareaField";
import Button from "../components/Button";
import { useI18n } from "vue-i18n";
import { ref } from "@vue/reactivity";
import useMessage from "../services/useMessage";

export default {
    name: "LetterShow",
    components: {
        Button,
        TextareaField,
        Letter,
        Page,
    },
    props: {
        messageId: String,
    },
    setup(props) {
        const { t } = useI18n();

        const showRaw = ref(false);
        const toggleMode = () => {
            showRaw.value = !showRaw.value;
        };

        const {
            message,
            isPending,
            isDestroyPending,
            isMissing,
            readableDuration,
            destroyMessage,
        } = useMessage(props.messageId);

        return {
            message,
            isPending,
            isDestroyPending,
            isMissing,
            readableDuration,
            showRaw,

            t,
            toggleMode,
            destroyMessage,
        };
    },
};
</script>

<style scoped></style>
