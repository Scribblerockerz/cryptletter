<template>
    <Page>
        <Letter striped spaced foot-separator>
            <TextareaField :placeholder="$t('messagePlaceholder')" />
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
            <Button primary>Encrypt message</Button>
        </div>
    </Page>
</template>

<script>
import Page from "../components/Page";
import Letter from "../components/Letter";
import TextareaField from "../components/TextareaField";
import Button from "../components/Button";

export default {
    name: "LetterNew",
    components: {
        Button,
        TextareaField,
        Letter,
        Page,
    },
};
</script>

<style scoped></style>
