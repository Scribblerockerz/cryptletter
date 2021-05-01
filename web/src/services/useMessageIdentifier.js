import { ref } from '@vue/reactivity';
import { onMounted, onUnmounted } from '@vue/runtime-core';
import MessageService from './MessageService';

function getMessagePartsFromUrl() {
    const [messageId, secret] = window.location.hash.substr(2).split('/');
    return {
        messageId: messageId || null,
        secret: secret || null,
    };
}

/**
 * Retrieve message id and setup the secret
 */
export default function useMessageIdentifier() {
    const messageId = ref(null);

    const locationChange = () => {
        const parts = getMessagePartsFromUrl();
        if (parts.secret === null) return;

        messageId.value = parts.messageId;
        MessageService.updateSecret(parts.secret);
    };

    onMounted(() => {
        locationChange();
        window.addEventListener('hashchange', locationChange);
    });

    onUnmounted(() => {
        window.removeEventListener('hashchange', locationChange);
    });

    return messageId;
}
