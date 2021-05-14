<template>
    <div v-if="files" class="attachment-list">
        <div v-for="(file, i) in files" :key="i" class="attachment-list__file">
            [{{ file.mimeType }}] {{ file.name }} ({{ file.encryptedSize }} /
            {{ file.originalSize }})
            <Button v-if="file.token" nano @click="$emit('downloadFile', file)">
                Download
            </Button>
            <Button v-else nano @click="$emit('removeFile', file)">
                Delete
            </Button>
        </div>
    </div>
</template>

<script>
import Button from "./Button";

export default {
    name: "AttachmentList",
    components: {
        Button,
    },
    props: {
        files: {
            type: Array,
            default: null,
        },
    },
    emits: ["removeFile"],
};
</script>

<style scoped>
.attachment-list {
}
</style>
