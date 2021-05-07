import { ref } from '@vue/reactivity';

const toasts = ref([]);

export default function useToaster() {
    const addToast = (message, type = 'info', delayMS = 5000) => {
        const toast = {
            id: Date.now(),
            message,
            type,
            delayMS,
        };

        toasts.value.push(toast);

        setTimeout(() => {
            toasts.value = toasts.value.filter((i) => i.id !== toast.id);
        }, delayMS);
    };

    return {
        toasts,
        addToast,
    };
}
