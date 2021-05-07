<template>
    <LetterShow v-if="messageId" :message-id="messageId" />
    <LetterNew v-else />
    <PasswordPrompt />
</template>

<script>
import { onMounted } from "@vue/runtime-core";
import "./assets/global.css";

import LetterNew from "./pages/LetterNew";
import LetterShow from "./pages/LetterShow";
import PasswordPrompt from "./pages/PasswordPrompt";

import useMessageIdentifier from "./services/useMessageIdentifier";

export default {
    name: "App",

    components: {
        LetterNew,
        LetterShow,
        PasswordPrompt,
    },

    setup() {
        const messageId = useMessageIdentifier();

        onMounted(() => {
            document.body.classList.add("is--ready");
            document.body.classList.remove("is--not-ready");
        });

        return {
            messageId,
        };
    },
};
</script>
