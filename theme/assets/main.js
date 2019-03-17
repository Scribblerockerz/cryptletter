require('./appearance.scss');

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
 * Get secret url by token, secret
 */
var getSecretUrl = function (token, secret) {
    return window.location.href.replace(/\/+$/, '')
            + '/' + token + '/#/' + secret;
}

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

    // reduce delay if message is long
    delay = text.length >= 500 ? 0 : delay;

    // skip animation if message is very long
    if (text.length < 1000) {

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
    }

    setTimeout(finishCallback, delay * i);
};

var checkRemainingTime = function(subjectSelector) {
    var timestamp = $(subjectSelector).data('available-until');

    if (!timestamp) { return false; }

    var activeUntil = new Date(timestamp);
    activeUntil && setInterval(function () {
        if (activeUntil < new Date()) {
          $('.page').remove();
          location.reload();
        }
    }, 1000);
}

$(document).ready(function () {

    var secret = generateSecretKey();

    // Select all text inside read only container
    $('body').on('focus', '.js--select-text', function (e) {
        $(this).select();
    });

    // Encrypt page
    $('body').on('click', '.js--encrypt-message', function (e) {

        var $encryptButton = $(this);

        // disable button while processing
        $encryptButton.attr('disabled', 'disabled');

        var $messageField = $('#message');
        var $urlField = $('#url');
        var $formStage = $('.js--stage-form');
        var $linkStage = $('.js--stage-url');
        var template = $('#template-result-format').html();

        var encryptedMessage = AES.encrypt($messageField.val().trim(), secret).toString();
        var delay = $('#delay').val();

        $('#selected-delay').text(delay);

        animateEncryptionOnText($messageField, function () {
            $.post('/', { message: encryptedMessage, delay: delay }, function (res) {
                if (res.token) {
                    var secureUrl = getSecretUrl(res.token, secret);
                    $urlField.html(template.replace('{minutes}', delay).replace('{url}', secureUrl));
                    $formStage.slideUp(300);
                    $linkStage.slideDown(300);
                    $urlField.select();
                    $encryptButton.fadeOut(300);
                }
            }).fail(function () {
                $('#error').slideDown(300);
            });
        });
    });


    // Decrypt
    if ($('.page--show').length) {
        var secret = window.location.hash.substr(2);
        var message = $('#encrypted-message').html();

        var encryptedMessage = AES.decrypt(message, secret).toString(encodingUTF8);
        if (encryptedMessage) {
            $('#message').text(encryptedMessage);
            $('#raw-message').val(encryptedMessage);
        }

        checkRemainingTime('#letter');

        //--- Stages switching

        $('body').on('click', '.js--toggle-stage', function () {
           var $el = $(this);
           var mode = $el.data('view');

            $('[data-view]')
                .hide()
                .filter('[data-view!="' + mode + '"]')
                .show();

            $('[data-stage]')
                .hide()
                .filter('[data-stage="' + mode + '"]')
                .show();
        });
    }

    // Handle destroy button
    $('body').on('click', '.js--destroy-message', function (e) {

        var $destroyButton = $(this).attr('disabled', 'disabled');
        var actionUrl = $destroyButton.data('action');

        if (actionUrl) {
          $.ajax({
            url: actionUrl,
            type: 'DELETE',
            success: function () {
              location.reload();
            }
          }).fail(function () {
              $('#error').slideDown(300);
          });
        }
    });

    // application is finsihed loading
    $('body')
        .addClass('is--ready')
        .removeClass('is--not-ready');
});
