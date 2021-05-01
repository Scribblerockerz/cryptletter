<template>
    <LetterShow v-if="messageId" :message-id="messageId" />
    <LetterNew v-else />
</template>

<script>
import "./assets/global.css";

import LetterNew from "./pages/LetterNew";
import LetterShow from "./pages/LetterShow";

import useMessageIdentifier from "./services/useMessageIdentifier";

export default {
    name: "App",

    components: {
        LetterNew,
        LetterShow,
    },

    setup() {
        const messageId = useMessageIdentifier();
        return {
            messageId,
        };
    },
};
</script>
