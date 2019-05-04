// require('./appearance.scss');
import './appearance.scss';

import AES from 'crypto-js/aes';
import encodingUTF8 from 'crypto-js/enc-utf8';
import sha512 from 'crypto-js/sha512';

/**
 * Generate secret key
 */
var generateSecretKey = function() {
    var crypto = window.crypto || window.msCrypto;
    var seed = Math.random();
    if (crypto) {
        seed = crypto.getRandomValues(new Uint8Array(2048 / 8)).toString();
    }

    return sha512(seed).toString();
};

/**
 * Get secret url by token, secret
 */
var getSecretUrl = function(token, secret) {
    const baseUrl = window.location.href.replace(/\/+$/, '');
    return `${baseUrl}/${token}/#/${secret}`;
};

/**
 * Helper function for replacing strings at index
 *
 * @param str
 * @param index
 * @param chr
 * @returns {*}
 */
var stringReplaceAt = function(str, index, chr) {
    if (index > str.length - 1) return str;
    return str.substr(0, index) + chr + str.substr(index + 1);
};

/**
 * Eyecandy — can be removed
 *
 * @param $elem
 * @param finishCallback
 * @param delay
 */
var animateEncryptionOnText = function($elem, finishCallback, delay) {
    delay = delay || 5;
    var text = $elem.value;
    var charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    // reduce delay if message is long
    delay = text.length >= 500 ? 0 : delay;
    // skip animation if message is very long
    if (text.length < 1000) {
        for (var i = 0; i < text.length; i++) {
            setTimeout(
                function(idx) {
                    var char = text.charAt(idx);
                    if (char !== '\n') {
                        // random character
                        var randChar = charset.charAt(Math.floor(Math.random() * charset.length));
                        text = stringReplaceAt(text, idx, randChar);
                    }
                    $elem.value = text;
                },
                delay * i,
                i,
            );
        }
    }
    setTimeout(finishCallback, delay * i);
};

var checkRemainingTime = function(subjectSelector) {
    const timestamp = document.querySelector(subjectSelector).dataset.availableUntil;
    if (!timestamp) return false;

    const activeUntil = new Date(parseInt(timestamp));
    const timeCheck = () => {
        const remainingMs = activeUntil - new Date();
        if (remainingMs <= 0) {
            const $page = document.querySelector('.page');

            while ($page.hasChildNodes()) {
                $page.removeChild($page.lastChild);
            }
        } else if (remainingMs > 0) {
            document.querySelector('#remaining-time').innerText = msToReadableDuration(remainingMs);
        }

        return true;
    };
    activeUntil && timeCheck() && setInterval(timeCheck, 1000);
};

/**
 * Transform miliseconds to a readable string
 *
 * @param {*} selector
 * @param {*} remainingMs
 */
var msToReadableDuration = function(remainingMs, skipSeconds = false) {
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
};

// TODO: implement fetch polyfil
// Implementation: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch
function postData(url = '', data = {}) {
    // Default options are marked with *
    return fetch(url, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, cors, *same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json',
            // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrer: 'no-referrer', // no-referrer, *client
        body: JSON.stringify(data), // body data type must match "Content-Type" header
    }).then(response => response.json()); // parses JSON response into native Javascript objects
}

// Delete Data
function deleteData(url = '') {
    // Default options are marked with *
    return fetch(url, {
        method: 'DELETE', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, cors, *same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
            'Content-Type': 'application/json',
        },
        redirect: 'follow', // manual, *follow, error
        referrer: 'no-referrer', // no-referrer, *client
        // body: '', // body data type must match "Content-Type" header
    }).then(response => response.json()); // parses JSON response into native Javascript objects
}

function fadeOut(el) {
    el.style.display = 'none';
}

function fadeIn(el) {
    el.style.display = '';
}

function slideUp(el) {
    el.style.display = 'none';
}

function slideDown(el) {
    el.style.display = '';
}

(function() {
    var secret = generateSecretKey();

    // Encrypt page
    document.querySelectorAll('.js--encrypt-message').forEach(node => {
        node.addEventListener('click', function(e) {
            const $encryptButton = e.target;

            // disable button while processing
            $encryptButton.disabled = true;

            const $messageField = document.getElementById('message');
            const $urlField = document.getElementById('url');
            const $formStage = document.querySelector('.js--stage-form');
            const $linkStage = document.querySelector('.js--stage-url');
            const template = document.getElementById('template-result-format').innerHTML;

            const encryptedMessage = AES.encrypt($messageField.value.trim(), secret).toString();
            const delay = parseInt(document.getElementById('delay').value);

            const readableDuration = msToReadableDuration(delay * 60 * 1000, true);

            document.getElementById('selected-delay').innerText = readableDuration;

            animateEncryptionOnText($messageField, function() {
                postData('/', { message: encryptedMessage, delay: delay })
                    .then(res => {
                        if (res.token) {
                            var secureUrl = getSecretUrl(res.token, secret);
                            $urlField.innerHTML = template
                                .replace('{duration}', readableDuration)
                                .replace('{url}', secureUrl);
                            // TODO: implement slideUp
                            slideUp($formStage);
                            // TODO: implement slideDown
                            slideDown($linkStage);
                            $urlField.select();
                            // TODO: implement fadeOut
                            fadeOut($encryptButton);
                        }
                    })
                    .catch(function() {
                        // TODO: implement slideDown by class
                        slideDown(document.getElementById('error'));
                    });
            });
        });
    });

    // Decrypt message
    if (document.querySelector('.page--show')) {
        const secret = window.location.hash.substr(2);
        const message = document.getElementById('encrypted-message').innerHTML;

        const encryptedMessage = AES.decrypt(message, secret).toString(encodingUTF8);
        if (encryptedMessage) {
            document.getElementById('message').innerText = encryptedMessage;
            document.getElementById('raw-message').innerHTML = encryptedMessage;
        }

        checkRemainingTime('#letter');

        //--- Stages switching
        document.querySelectorAll('.js--toggle-stage').forEach(node => {
            node.addEventListener('click', function(e) {
                const $el = e.target;
                const mode = $el.dataset.view;

                document.querySelectorAll('[data-view]').forEach(node => {
                    const attribute = node.attributes['data-view'];

                    if (attribute && attribute.nodeValue !== mode) {
                        node.style.display = 'inline-block';
                    } else {
                        node.style.display = 'none';
                    }
                });

                document.querySelectorAll('[data-stage]').forEach(node => {
                    const attribute = node.attributes['data-stage'];

                    if (attribute && attribute.nodeValue === mode) {
                        node.style.display = 'block';
                    } else {
                        node.style.display = 'none';
                    }
                });
            });
        });
    }

    // Destroy message
    document.querySelectorAll('.js--destroy-message').forEach(node => {
        node.addEventListener('click', function(e) {
            const $destroyButton = e.target;
            $destroyButton.disabled = true;
            const actionUrl = $destroyButton.dataset.action;

            if (actionUrl) {
                deleteData(actionUrl)
                    .then(() => {
                        window.location.reload();
                    })
                    .catch(() => {
                        // TODO: implement slideDown
                        // document.getElementById('error').slideDown(300);
                        slideDown(document.getElementById('error'));
                    });
            }
        });
    });

    // Select all text inside read only container
    document.querySelectorAll('.js--select-text').forEach(node => {
        node.addEventListener('focus', function(e) {
            e.target.select();
        });
    });

    const $body = document.getElementsByTagName('body')[0];
    $body.classList.add('is--ready');
    $body.classList.remove('is--not-ready');
})();
