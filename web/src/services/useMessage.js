import { ref } from '@vue/reactivity';
import MessageService from './MessageService';
import { msToReadableDuration } from './utils';

const REMAINING_TIME_INTERVAL_MS = 1000;

/**
 * Request message information
 *
 * @param {string} messageId
 */
export default function useMessage(messageId = null) {
    let activeUntilTimestamp = null;
    let remainingMs = null;
    let remainingTimeInterval = null;

    const isPending = ref(true);
    const isDestroyPending = ref(false);
    const isMissing = ref(false);
    const message = ref(null);
    const attachments = ref([]);
    const readableDuration = ref(null);

    const checkRemainingTime = () => {
        remainingMs = activeUntilTimestamp - Date.now();

        if (remainingMs > 0) {
            readableDuration.value = msToReadableDuration(remainingMs);
            return;
        }

        // self destruction
        const body = document.body;
        while (body.hasChildNodes()) {
            body.removeChild(body.lastChild);
        }
    };

    MessageService.getMessageData(messageId)
        .then((data) => {
            message.value = data.message;
            attachments.value = data.attachments;
            activeUntilTimestamp = parseInt(data.activeUntilTimestamp);
            checkRemainingTime();
            isPending.value = false;

            remainingTimeInterval && clearInterval(remainingTimeInterval);
            remainingTimeInterval = setInterval(
                checkRemainingTime,
                REMAINING_TIME_INTERVAL_MS
            );
        })
        .catch((err) => {
            if (err?.response?.status !== 404) {
                console.log(err, err.message);
                return;
            }

            isMissing.value = true;
            isPending.value = false;
        });

    const destroyMessage = async () => {
        isDestroyPending.value = true;
        try {
            await MessageService.destory(messageId);
        } catch (err) {
            // noop
        }
        remainingTimeInterval && clearInterval(remainingTimeInterval);
        message.value = null;
        isMissing.value = true;
        isDestroyPending.false = true;
    };

    return {
        isPending,
        isDestroyPending,
        isMissing,
        message,
        attachments,
        readableDuration,
        destroyMessage,
    };
}
