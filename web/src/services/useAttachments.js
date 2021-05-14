import { ref } from '@vue/reactivity';
import { watchEffect } from '@vue/runtime-core';
import MessageService from './MessageService';
import { encryptFile, decryptString } from './utils/fileEncryption';

function createFileInput() {
    const node = document.createElement('input');
    node.type = 'file';
    node.style.opacity = '0';
    node.tabIndex = -1;
    node.style.position = 'fixed';
    node.style.top = '-9999em';
    node.multiple = true;
    document.body.append(node);

    return node;
}

export function useAttachmentsWriter(encryptionKey) {
    const $fileInput = createFileInput();

    const files = ref([]);

    $fileInput.addEventListener('change', (event) => {
        if (event.target.files.length === 0) return;

        const currentFileNames = files.value.map((f) => f.name);

        Array.from(event.target.files).forEach((file) => {
            // Skip files with the same name
            if (currentFileNames.includes(file.name)) return;

            const encryptedFile = ref(null);

            encryptFile(file, encryptionKey).then((file) => {
                encryptedFile.value = file;
            });
            files.value = [...files.value, encryptedFile];
        });

        // Reset actual file input
        $fileInput.value = '';
    });

    // Trigger file selection
    function selectFiles() {
        $fileInput.click();
    }

    // Remove file from the list of managed files
    function removeFile(file) {
        if (!file) return;
        files.value = files.value.filter((f) => f.value.name !== file.name);
    }

    return {
        files,

        selectFiles,
        removeFile,
    };
}

export function useAttachmentsReader(
    encryptionKey,
    messageId,
    encryptedAttachments
) {
    const files = ref([]);
    watchEffect(() => {
        files.value = encryptedAttachments.value.map((encryptedAttachment) => {
            const file = ref(null);
            file.value = {
                name: decryptString(encryptedAttachment.name, encryptionKey),
                mimeType: decryptString(
                    encryptedAttachment.mimeType,
                    encryptionKey
                ),
                size: parseInt(
                    decryptString(encryptedAttachment.size, encryptionKey)
                ),
                token: encryptedAttachment.token,
            };
            return file;
        });

        console.log(files);
    });

    console.log(files);

    async function downloadFile(fileInfo) {
        //
        console.log('Download', {
            token: messageId,
            attachmentToken: fileInfo.token,
        });

        const file = await MessageService.downloadAttachment(
            messageId,
            fileInfo.token,
            fileInfo.mimeType,
            fileInfo.name
        );

        console.log('Actual file', file);
    }

    return {
        files,

        downloadFile,
    };
}
