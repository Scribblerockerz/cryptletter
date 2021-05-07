<template>
    <div class="toaster">
        <Alert
            v-for="(toast, i) in toasts"
            :key="i"
            :type="toast.type"
            class="toaster__toast"
            >{{ toast.message }}</Alert
        >
    </div>
</template>

<script>
import Alert from "./Alert";
import useToaster from "../services/useToaster";

export default {
    name: "Toaster",
    components: {
        Alert,
    },
    setup() {
        const { toasts } = useToaster();

        return {
            toasts,
        };
    },
};
</script>

<style scoped>
.toaster__toast {
    margin-bottom: 20px;
}
</style>
