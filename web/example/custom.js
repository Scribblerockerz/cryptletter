// Additional configuration
window.cryptletterExtended = {
    // change active locale
    activeLocale: 'de',

    // add additional locales or update existing ones
    locales: {
        de: {
            slogan: 'Encrypted self-destructing messages.',
            messagePlaceholder:
                'Diesen Text entfernen und die entsprechende Nachricht hier eintragen.\n\nDie Nachricht wird verschlüsselt übertragen und ist nach der Erstellung nur einmalig aufrufbar.',
            lifetimeHint:
                'Dieser Cryptletter wird {selectDuration} nach dem ersten Aufruf unwiderruflich gelöscht.',
            lifetimeHintAfter:
                'Dieser Link ist nach dem erstmaligen Aufruf {duration} gültig.',
            estimatedLifetimeHint:
                'Dieser Cryptletter wird in {duration} unwiderruflich gelöscht.',
            encryptMessageLabel: 'Nachricht verschlüsseln',
            destroyMessageLabel: 'Nachricht jetzt löschen',
            messageResult:
                'Der folgende Link ist nach dem ersten Aufruf nur {duration} verfügbar.\n\n\n{url}\n',
            messageEmpty: 'Looks like there was no message provided.',
            decryptingMessage: 'Decrypting message, please wait ...',
            missingMessageHeadline:
                'Dieser Cryptletter ist nicht mehr verfügbar...',
            missingMessageText:
                'Wenn Sie hier stattdessen eine verschlüsselte Nachricht erwarten, kontaktieren Sie bitte Ihren Ansprechpartner.',
            viewRawLabel: 'view raw',
            viewTextLabel: 'view text',
            addAttachmentsLabel: "Anhang hinzufügen",
            deleteAttachmentLabel: "löschen",
            downloadAttachmentLabel: "download",
            encryptionInProgressFileNameLabel: "Wird verschlüsselt ...",
            encryptionInProgressFileSizeLabel: "Größe wird geschätzt ...",
            encryptionInProgressAttachmentLabel: "Bitte warten",
            passwordProtection: {
                headline: "Postage Stamp Protection",
                description: "Diese Instanz erfordert ein Passwort um einen Cryptletter zu erstellen.",
                passwordLabel: "Passwort",
                confirmLabel: "Bestätigen",
                cancelLabel: "Abbrechen"
            },
            errors: {
                publishing_failed: "Wir konnten leider Ihren Cryptletter nicht ausstellen."
            }
        },
    },
};
