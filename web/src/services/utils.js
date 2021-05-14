import sha512 from 'crypto-js/sha512';

/**
 * Generate secret key
 */
export function generateSecretKey() {
    var crypto = window.crypto || window.msCrypto;
    var seed = Math.random();
    if (crypto) {
        seed = crypto.getRandomValues(new Uint8Array(2048 / 8)).toString();
    }

    return sha512(seed).toString();
}

/**
 * Transform miliseconds to a readable string
 *
 * @param {*} selector
 * @param {*} remainingMs
 */
export function msToReadableDuration(remainingMs, skipSeconds = false) {
    if (remainingMs <= 0) return null;

    const s = remainingMs / 1000;
    const m = s / 60;
    const h = m / 60;
    const d = h / 24;

    const seconds = Math.floor(s % 60);
    const minutes = Math.floor(m % 60);
    const hours = Math.floor(h % 60);
    const days = Math.floor(d);

    const parts = [];

    Math.floor(d) > 0 && parts.push(`${days}d`);
    Math.floor(h) > 0 && parts.push(`${hours}h`);
    Math.floor(m) > 0 && parts.push(`${minutes}min`);
    (skipSeconds && s > 60) || (s >= 0 && parts.push(`${seconds}s`));

    return parts.join(' ');
}

/**
 * Helper function for replacing strings at index
 *
 * @param str
 * @param index
 * @param chr
 * @returns {*}
 */
function stringReplaceAt(str, index, chr) {
    if (index > str.length - 1) return str;
    return str.substr(0, index) + chr + str.substr(index + 1);
}

/**
 * Eyecandy — can be removed
 *
 * @param $elem
 * @param finishCallback
 * @param delay
 */
export async function animateEncryptionOnText(
    text,
    runtimeCallback,
    delay = 5
) {
    return new Promise((resolve) => {
        // skip animation if message is very long
        if (text.length >= 1000) {
            resolve();
        }

        const charset =
            'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

        // reduce delay if message is long
        delay = text.length >= 500 ? 0 : delay;

        for (var i = 0; i < text.length; i++) {
            setTimeout(
                function (idx) {
                    var char = text.charAt(idx);
                    if (char !== '\n') {
                        // random character
                        var randChar = charset.charAt(
                            Math.floor(Math.random() * charset.length)
                        );
                        text = stringReplaceAt(text, idx, randChar);
                    }
                    runtimeCallback(text);
                },
                delay * i,
                i
            );
        }

        setTimeout(resolve, delay * i);
    });
}

/**
 * Copied from some dude on the internet
 *
 * @param {*} bytes
 * @param {*} si
 */
export function humanFileSize(bytes, si) {
    let thresh = si ? 1000 : 1024;
    if (Math.abs(bytes) < thresh) {
        return `${bytes} B`;
    }
    let units = si
        ? ['kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
        : ['KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];
    let u = -1;
    do {
        bytes /= thresh;
        ++u;
    } while (Math.abs(bytes) >= thresh && u < units.length - 1);
    return `${bytes.toFixed(1)} ${units[u]}`;
}
