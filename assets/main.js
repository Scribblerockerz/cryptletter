require('./appearance.css');

var AES = require("crypto-js/aes");
var encodingUTF8 = require("crypto-js/enc-utf8");
var sha512 = require("crypto-js/sha256");

var $ = require('jquery');

/**
 * Generate secret key
 */
var generateSecretKey = function () {
    var crypto = window.crypto || window.msCrypto;
    var seed = Math.random();

    if (crypto) {
        seed = crypto.getRandomValues(new Uint8Array(2048 / 8)).toString();
    }

    return sha512(seed).toString();
};

/**
 * Helper function for replacing strings at index
 *
 * @param str
 * @param index
 * @param chr
 * @returns {*}
 */
var stringReplaceAt = function (str, index, chr) {
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
var animateEncryptionOnText = function ($elem, finishCallback, delay) {
    delay = delay || 5;

    var text = $elem.val();
    var charset = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

    for (var i=0; i<text.length;i++) {
        setTimeout(function(idx) {
            var char = text.charAt(idx);
            if (char !== '\n') {
                // random character
                var randChar = charset.charAt(Math.floor(Math.random() * charset.length));
                text = stringReplaceAt(text, idx, randChar);
            }
            $elem.val(text);
        }, delay * i, i);
    }

    setTimeout(finishCallback, delay * i);
};

$(document).ready(function () {

    // Encrypt page
    var $form = $('.form--message');
    if ($form.length) {
        var secret = generateSecretKey();

        $form.on('submit', function (e) {
            e.preventDefault();
            e.stopPropagation();
            $('#btn--encrypt').attr('disabled', 'disabled');


            var $field = $('#message');
            var encryptedMessage = AES.encrypt($field.val(), secret).toString();
            var delay = $('#destructionDelay').val();

            $('#selected-delay').text(delay);

            animateEncryptionOnText($field, function () {
                $.post('/', { message: encryptedMessage, delay: delay }, function (res) {
                    if (res.success && res.token) {
                        $('#url').val(window.location.href.replace(/\/+$/, '') + '/' + res.token + '/#/' + secret);
                        $('.form--message').slideUp(300);
                        $('#secret').slideDown(300);
                        $('#url').select();
                    }
                });
            });

            return false;
        });
    }

    // Decrypt
    if ($('.page--show').length) {
        var secret = window.location.hash.substr(2);
        var message = $('#encrypted-message').html();

        var encryptedMessage = AES.decrypt(message, secret).toString(encodingUTF8);
        if (encryptedMessage) {
            $('#message').text(encryptedMessage);
            $('#raw-message').val(encryptedMessage);
        }

        $('body').on('click', '.pane__button', function () {
           var $el = $(this);
           var mode = $el.data('view');

            $('.pane__button[data-view]')
                .hide()
                .filter('[data-view!="' + mode + '"]')
                .show();

            $('.secret__message[data-stage]')
                .hide()
                .filter('[data-stage="' + mode + '"]')
                .show();

            $('#raw-message').select();
        });
    }

    if ($('#form-destroy-message')) {
      $('body').on('submit', '#form-destroy-message', function (e) {
        e.preventDefault();
        e.stopPropagation();

        if (this.action) {
          $.ajax({
            url: this.action,
            type: 'DELETE',
            success: function (res) {
              location.reload();
            }
          });
        }
      });
    }
});
