<template>
    <Page v-if="!isPending && isMissing" :is-unavailable="isMissing">
        <h2>{{ t("missingMessageHeadline") }}</h2>
        <p>{{ t("missingMessageText") }}</p>
    </Page>
    <Page v-if="!isPending && !isMissing">
        <Letter striped spaced foot-separator>
            <template v-slot:header>
                <Button nano @click="toggleMode">
                    {{ showRaw ? t("viewRawLabel") : t("viewTextLabel") }}
                </Button>
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
            <template v-slot:attachment>
                <AttachmentList :files="files" @downloadFile="downloadFile" />
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
import AttachmentList from "../components/AttachmentList";
import { useAttachmentsReader } from "../services/useAttachments";
import MessageService from "../services/MessageService";

export default {
    name: "LetterShow",
    components: {
        Button,
        TextareaField,
        Letter,
        Page,
        AttachmentList,
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
            attachments,
            isPending,
            isDestroyPending,
            isMissing,
            readableDuration,
            destroyMessage,
        } = useMessage(props.messageId);

        const { files, downloadFile } = useAttachmentsReader(
            MessageService.getSecret(),
            props.messageId,
            attachments
        );

        return {
            message,
            isPending,
            isDestroyPending,
            isMissing,
            readableDuration,
            showRaw,
            files,

            t,
            toggleMode,
            destroyMessage,
            downloadFile,
        };
    },
};
</script>

<style scoped></style>
