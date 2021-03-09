<template>
    <Page>
        <Letter striped spaced foot-separator>
            <template v-slot:header>
                <Button nano>view raw</Button>
                <Button nano>view text</Button>
            </template>
            <TextareaField :placeholder="$t('messagePlaceholder')" />
            <!--            <pre-->
            <!--                style="display: block"-->
            <!--            ><code id="message">Decrypting message, please wait ...</code></pre>-->

            <template v-slot:footer>
                <i18n-t keypath="lifetimeHint" tag="span">
                    <template #selectDuration>
                        <select id="delay" name="delay">
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
        <div class="u--center">
            <Button primary>{{ $t("encryptMessageLabel") }}</Button>
        </div>
    </Page>
</template>

<script>
import Page from "../components/Page";
import Letter from "../components/Letter";
import TextareaField from "../components/TextareaField";
import Button from "../components/Button";

export default {
    name: "LetterShow",
    components: {
        Button,
        TextareaField,
        Letter,
        Page,
    },
};
</script>

<style scoped></style>
