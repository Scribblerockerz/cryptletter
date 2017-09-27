# Cryptletter

Self-hosted micro-service for encrypted self-destructing messages.

**Current state:** *proof of concept*

---

Sending plaintext passwords unencrypted through the internet highway isn't just risky, it's ridiculous.
This project aims to make this process a bit more secure.

Retain control over the data which is send out and prevent living-security-issues laying around in the users inbox.

## Features

### Self-hosted

Grab it. Use it.

### Client side AES-256 encryption

Messages are encrypted on the client side with the amazing [crypto-js](https://www.npmjs.com/package/crypto-js) library.

The key is appended as a hash, so it never hits the server. (*In case of a hack on the micro-service, only encrypted garbage is captured*.)

### Messages with a TTL (Time To Live)

Decide how long this message can survive after the client opens it.

### Restricted message access

Messages are restricted to the client's IP address at the first opening. Messages posted through facebook or other link checking tools prevent opening of the message.

---

## Setup

Requirements: mysql/node
...

## Build

Build your own executable version with [pkg](https://www.npmjs.com/package/pkg).

Install `pkg` globally and run `npm run build`.

**...WIP...**
