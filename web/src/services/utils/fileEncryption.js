import CryptoJS from 'crypto-js';
import encodingUTF8 from 'crypto-js/enc-utf8';

export async function encryptFile(file, encryptionKey) {
    return new Promise((resolve) => {
        const mimeType = file.type;

        const reader = new FileReader();
        reader.onload = (e) => {
            const fileData = e.target.result;

            // Create a WordArray out of ArrayBuffer
            let wordBuffer = CryptoJS.lib.WordArray.create(fileData);

            // Encrypt data
            const encryptedData = encrypt(wordBuffer, encryptionKey);

            // Transform data to plain string (used for transfer to server)
            const encryptedDataAsString = encryptedData.toString();
            const encryptedName = encryptString(file.name, encryptionKey);
            const encryptedMimeType = encryptString(mimeType, encryptionKey);
            const encryptedSize = encryptString(`${file.size}`, encryptionKey);

            resolve({
                name: file.name,
                encryptedName,
                mimeType: mimeType,
                encryptedMimeType,
                size: file.size,
                encryptedSize,
                data: encryptedDataAsString,
            });
        };
        reader.readAsArrayBuffer(file);
    });
}

export function decryptFileData(encryptedFileData, encryptionKey) {
    // Decrypt data
    const decryptedData = decrypt(encryptedFileData, encryptionKey);

    // Convert WordArray to Uint8Array
    const uint8buffer = wordArrayToUint8Array(decryptedData);

    return uint8buffer;
}

// https://stackoverflow.com/a/35536933/1704139
function encrypt(data, key) {
    return CryptoJS.AES.encrypt(data, key);
}

function encryptString(data, key) {
    return encrypt(data, key).toString();
}

function decrypt(encryptedData, key) {
    return CryptoJS.AES.decrypt(encryptedData, key);
}

export function decryptString(encryptedData, key) {
    return decrypt(encryptedData, key).toString(encodingUTF8);
}

/*
 * Converts a cryptjs WordArray to native Uint8Array
 * https://github.com/brix/crypto-js/issues/274
 */
function wordArrayToUint8Array(wordArray) {
    const l = wordArray.sigBytes;
    const words = wordArray.words;
    const result = new Uint8Array(l);
    var i = 0 /*dst*/,
        j = 0; /*src*/

    // eslint-disable-next-line no-constant-condition
    while (true) {
        // here i is a multiple of 4
        if (i == l) break;
        var w = words[j++];
        result[i++] = (w & 0xff000000) >>> 24;
        if (i == l) break;
        result[i++] = (w & 0x00ff0000) >>> 16;
        if (i == l) break;
        result[i++] = (w & 0x0000ff00) >>> 8;
        if (i == l) break;
        result[i++] = w & 0x000000ff;
    }

    return result;
}
