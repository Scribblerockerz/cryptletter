<template>
    <Page>
        <Letter striped spaced foot-separator>
            <TextareaField
                v-show="!isSubmitted"
                v-model="message"
                :placeholder="t('messagePlaceholder')"
            />
            <TextareaField
                v-show="isSubmitted"
                :modelValue="resultMessage"
                :rows="11"
                ref="resultInput"
                url-holder
                readonly
                focus-selection
            />
            <template v-slot:footer>
                <i18n-t
                    v-if="isSubmitted"
                    keypath="lifetimeHintAfter"
                    tag="span"
                >
                    <template #duration>{{
                        durationInWords || "..."
                    }}</template>
                </i18n-t>
                <i18n-t v-else keypath="lifetimeHint" tag="span">
                    <template #selectDuration>
                        <select v-model="delay" id="delay" name="delay">
                            <option value="15">15min</option>
                            <option value="30">30min</option>
                            <option value="60">1h</option>
                            <option value="120">2h</option>
                            <option value="1440">24h</option>
                        </select>
                    </template>
                </i18n-t>
            </template>
        </Letter>
        <div v-if="!isSubmitted" class="u--center">
            <Button :disabled="isPending" primary @click="submit">{{
                t("encryptMessageLabel")
            }}</Button>
        </div>
    </Page>
</template>

<script>
import Page from "../components/Page";
import Letter from "../components/Letter";
import TextareaField from "../components/TextareaField";
import Button from "../components/Button";
import { ref, computed } from "@vue/reactivity";
import MessaageService from "../services/MessageService";
import {
    animateEncryptionOnText,
    msToReadableDuration,
} from "../services/utils";
import { useI18n } from "vue-i18n";

export default {
    name: "LetterNew",
    components: {
        Button,
        TextareaField,
        Letter,
        Page,
    },
    setup() {
        const { t } = useI18n();

        const message = ref("");
        const delay = ref(15);
        const isPending = ref(false);
        const isSubmitted = ref(false);
        const url = ref(null);
        const durationInWords = ref(null);
        const resultInput = ref(null);

        const resultMessage = computed(() => {
            if (!url.value) return "";
            return t("messageResult", {
                url: url.value,
                duration: durationInWords.value,
            });
        });

        async function submit() {
            const d = parseInt(delay.value);
            isPending.value = true;
            const rawMessage = message.value;

            await animateEncryptionOnText(message.value, (nextText) => {
                message.value = nextText;
            });

            url.value = await MessaageService.publish(rawMessage, d);
            durationInWords.value = msToReadableDuration(d * 60 * 1000, true);
            isSubmitted.value = true;
            setTimeout(() => {
                resultInput.value.$el.focus();
            }, 100);
        }

        return {
            message,
            delay,
            isPending,
            isSubmitted,
            url,
            durationInWords,
            resultMessage,
            resultInput,

            submit,
            t,
        };
    },
};
</script>

<style scoped></style>
