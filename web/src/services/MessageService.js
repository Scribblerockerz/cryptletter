import axios from 'axios';
import AES from 'crypto-js/aes';
import encodingUTF8 from 'crypto-js/enc-utf8';
import { generateSecretKey } from './utils';
import { requestPasswordPrompt } from '../pages/PasswordPrompt';
import { decryptFileData } from './utils/fileEncryption';

const BASE_URL = process.env.VUE_APP_API_BASE_URL || '';

class MessageService {
    #secret;
    constructor() {
        this.updateSecret(null);
    }

    /**
     * Update secret
     * @param {string|null} secret
     */
    updateSecret(secret) {
        if (secret === null) {
            this.#secret = generateSecretKey();
            return;
        }

        this.#secret = secret;
    }

    getSecret() {
        return this.#secret;
    }

    /**
     *
     * @param {string} message
     * @param {number} delayInMinutes
     * @param {Array} files
     */
    async publish(message, delayInMinutes, files = []) {
        const encryptedMessage = AES.encrypt(
            message.trim(),
            this.#secret
        ).toString();

        let res;
        const payload = {
            message: encryptedMessage,
            delay: delayInMinutes,
            attachments: files.map((file) => {
                return {
                    name: file.encryptedName,
                    mimeType: file.encryptedMimeType,
                    size: file.encryptedSize,
                    data: file.data,
                };
            }),
        };

        try {
            res = await axios.post(`${BASE_URL}/api/`, payload);
        } catch (err) {
            if (err.response?.status === 401) {
                const creationRestrictionPassword = await requestPasswordPrompt();

                res = await axios.post(`${BASE_URL}/api/`, {
                    ...payload,
                    creationRestrictionPassword,
                });
            } else {
                throw err;
            }
        }

        const secureUrl = this.getSecretUrl(res.data.token, this.#secret);

        return secureUrl;
    }

    async downloadAttachment(messageId, attachmentId, mimeType, fileName) {
        const res = await axios.get(
            `${BASE_URL}/api/${messageId}/attachment/${attachmentId}`
        );

        const uint8buffer = decryptFileData(res.data, this.#secret);

        const url = window.URL.createObjectURL(
            new Blob([uint8buffer], { type: mimeType })
        );
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', fileName);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }

    /**
     * @param {string} message
     */
    async getMessageData(messageId) {
        const res = await axios.get(`${BASE_URL}/api/${messageId}/`);

        return {
            ...res.data,
            message: AES.decrypt(res.data.message, this.#secret).toString(
                encodingUTF8
            ),
        };
    }

    /**
     *
     * @param {string} message
     * @param {number} delayInMinutes
     */
    async destory(messageId) {
        await axios.delete(`${BASE_URL}/api/${messageId}/`);
    }

    /**
     * Get secret url
     * @param {string} token
     * @returns {string}
     */
    getSecretUrl(token) {
        const baseUrl = window.location.href.replace(/\/+$/, '');
        return `${baseUrl}/#/${token}/${this.#secret}`;
    }
}

export default new MessageService();
