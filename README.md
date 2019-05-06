<p align="center">
<img src="https://user-images.githubusercontent.com/1336156/31407710-267e2124-ae06-11e7-8a48-4e7dc3547422.png" />
<br>
</p>

## Introduction

Sending plaintext passwords unencrypted through the internet highway isn't just risky, it's ridiculous.
This project aims to make this process a bit more secure.

Usually an email inbox of a regular user contains more plaintext passwords than emails from rich african princes.

Retain control over the data which is send out, and prevent living-security-issues laying around in the users inbox.

## Demo

![Screen capture of the application](https://user-images.githubusercontent.com/1336156/32858885-31374bc8-ca4d-11e7-9b1a-b4a67769a241.gif)

## Features

- **Self-hosted**  
  Grab it. Use it.

- **Client side AES-256 encryption**  
  Messages are encrypted on the client side with the amazing [crypto-js](https://www.npmjs.com/package/crypto-js) library.

  The key is appended as a hash, so it never hits the server. (*In case of a hack on the micro-service, only encrypted garbage is captured*.)

- **Messages with a TTL**  
  Decide how long this message can survive after the client opens it.

- **Restricted message access**  
  Messages are restricted to the client's IP address at the first opening. Messages posted through facebook or other link checking tools prevent opening of the message.

---
## Quick usage

Create a `docker-compose.yml` with the following contents and run `docker-compose up`.

```yaml
# docker-compose.yml
version: '3'

services:
    redis:
        image: 'redis:alpine'
        ports:
            - '6379:6379'
    app:
        image: 'scribblerockerz/cryptletter:latest'
        command: './cryptletter -vvv'
        ports:
            - '8080:8080'
        links:
            - redis
        environment:
            - 'APP_DATABASE_ADDRESS=redis:6379'
```


## Requirements
This microservice should be run via docker. If you prefer to run it standalone, check the releases page for the latest executable.


## Changing Templates

The application can be modified by replacing templates and adding custom stylesheets.

##### Template override

Add a new template directory and configure it with the env variable `APP_APP_TEMPLATESDIR` or in the toml file.
```toml
[app]
templatesDir = "./my-theme/template-override"

```
Place all the templates in the same structure to be overwriten.

##### Additional Assets

Add a new assets directory and configure it with the env variable `APP_APP_ASSETSDIR` or in the toml file.
```toml
[app]
assetsDir = "./my-theme/assets"
```

All the assets will be mounted under `/static/`. Default assets are served under `/s/`.

```html
<link rel="stylesheet" href="/static/custom-theme.css">
```


## Build

Run `./build.sh` and get your executable.
