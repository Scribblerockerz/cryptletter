import { ref } from '@vue/reactivity';
import { encryptFileInput } from './utils/fileEncryption';

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

export default function useAttachments(encryptionKey) {
    const $fileInput = createFileInput();

    const files = ref([]);

    $fileInput.addEventListener('change', (event) => {
        if (event.target.files.length === 0) return;

        const currentFileNames = files.value.map((f) => f.name);

        Array.from(event.target.files).forEach(async (file) => {
            // Skip files with the same name
            if (currentFileNames.includes(file.name)) return;

            const encryptedFile = await encryptFileInput(file, encryptionKey);
            files.value = [...files.value, encryptedFile];
        });

        // Reset actual file input
        $fileInput.value.value = '';
    });

    // Trigger file selection
    function selectFiles() {
        $fileInput.click();
    }

    // Remove file from the list of managed files
    function removeFile(file) {
        if (!file) return;
        files.value = files.value.filter((f) => f.name !== file.name);
    }

    return {
        files,

        selectFiles,
        removeFile,
    };
}
