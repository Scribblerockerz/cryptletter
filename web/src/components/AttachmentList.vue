<template>
    <div v-if="files.length" class="attachment-list">
        <div v-for="(file, i) in files" :key="i" class="attachment-list__file">
            <FileIcon :mimeType="file.mimeType" />
            <div class="attachment-list__file-info">
                <span class="attachment-list__file-name">{{ file.name }}</span>
                <div class="attachment-list__file-size">
                    {{ humanFileSize(file.size, true) }}
                </div>
            </div>
            <Button v-if="file.token" nano @click="$emit('downloadFile', file)">
                {{ t("downloadAttachmentLabel") }}
            </Button>
            <Button v-else nano @click="$emit('removeFile', file)">
                {{ t("deleteAttachmentLabel") }}
            </Button>
        </div>
    </div>
</template>

<script>
import { humanFileSize } from "../services/utils";
import Button from "./Button";
import FileIcon from "./FileIcon";
import { useI18n } from "vue-i18n";

export default {
    name: "AttachmentList",
    components: {
        Button,
        FileIcon,
    },
    props: {
        files: {
            type: Array,
            default: null,
        },
    },
    emits: ["removeFile"],
    setup() {
        const { t } = useI18n();

        return {
            humanFileSize,
            t,
        };
    },
};
</script>

<style scoped>
.attachment-list {
    margin: 0 10px;
    background-color: #d9dce4;
    box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.2);
    position: relative;
    padding: 0 20px;
}

@media screen and (min-width: 769px) {
    .attachment-list {
        margin: 0 50px;
    }
}

.attachment-list::before,
.attachment-list::after {
    content: "";
    display: block;
    position: absolute;
    bottom: -10px;
    width: 0;
    height: 0;
    border-style: solid;
    border-color: #a6acbc transparent transparent transparent;
}

.attachment-list::before {
    right: 30px;
    border-width: 10px 2px 0 100px;
}

.attachment-list::after {
    left: 30px;
    border-width: 10px 100px 0 2px;
}

@media screen and (min-width: 769px) {
    .attachment-list::before {
        right: 30px;
        border-width: 10px 2px 0 300px;
    }

    .attachment-list::after {
        left: 30px;
        border-width: 10px 300px 0 2px;
    }
}

.attachment-list__file {
    color: #000;
    padding: 10px 0;
    display: grid;
    grid-template-columns: auto 1fr auto;
    gap: 10px;
    align-items: center;
    line-height: 1;
}

.attachment-list__file + .attachment-list__file {
    border-top: 1px solid #c2c5d0;
}

.attachment-list__file-info {
    overflow: hidden;
    text-overflow: ellipsis;
}

.attachment-list__file-name {
    white-space: pre;
}

.attachment-list__file-size {
    opacity: 0.5;
    font-size: 0.8em;
    margin-top: 5px;
}
</style>
