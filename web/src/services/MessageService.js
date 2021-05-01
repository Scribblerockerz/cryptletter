import axios from 'axios';
import AES from 'crypto-js/aes';
import encodingUTF8 from 'crypto-js/enc-utf8';
import { generateSecretKey } from './utils';

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

    /**
     *
     * @param {string} message
     * @param {number} delayInMinutes
     */
    async publish(message, delayInMinutes) {
        const encryptedMessage = AES.encrypt(
            message.trim(),
            this.#secret
        ).toString();

        const res = await axios.post(`${BASE_URL}/api/`, {
            message: encryptedMessage,
            delay: delayInMinutes,
        });

        const secureUrl = this.getSecretUrl(res.data.token, this.#secret);

        return secureUrl;
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
