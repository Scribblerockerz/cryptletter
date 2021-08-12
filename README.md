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
        ports:
            - '8080:8080'
        links:
            - redis
        environment:
            - 'REDIS__ADDRESS=redis:6379'
            - 'APP__LOG_LEVEL=4'
```

## Requirements
This microservice requires redis to work and can be run via docker or standalone executable.

## Configuration

Configuration can be provided via configuration yaml or env variables.

You can run `cryptletter config:init` to generate a fresh `cryptletter.yml` in your working directory.
You can also specify the config file by providing it as an argument to the executable:

```bash
$ cryptletter --config ../your/own/path/you-name-it.yml
```


```yml
# cryptletter.yml
app:
  # How long should the message survive, without getting opened? (minutes)
  default_message_ttl: 43830
  # LOUDER > quieter
  log_level: 4
  # Current env, use "dev" to disable cors for local development
  env: prod
  
  # Serving config
  server:
    port: 8080

  # Restrict creation of new letters with a password (good enough to lockout the public)
  creation_protection_password: ""
    
  # Inject custom css and custom js configuration
  additional:
    css: './web/example/additional.css'
    js: './web/example/custom.js'
    
  attachments:
    # Files must be removed if the message reached it's TTL and is no longer reachable 
    cleanup_schedule: * * * * *
    # Supported driver: s3, local or "" to disable attachment support
    driver: local
    # Directory for uploaded attachments
    storage_path: cryptletter-uploads

# Redis config
redis:
  address: 127.0.0.1:6379
  database: 0
  password: ""

# S3 configuration for attachment.driver: s3
s3:
  access_id: minioadmin
  access_secret: minioadmin
  bucket_name: cryptletter-attachments
  bucket_region: eu-central-1
  endpoint: http://127.0.0.1:9000
  secure: true

```

Environment variables can be used with `__` as the replacement for dot notation.

```
$ APP__LOG_LEVEL=0 cryptletter
```


## Customization

This microservice is designed to work as it is. It comes with an embedded version of the frontend app (thanks to [go:embed](https://golang.org/pkg/embed/)).

It's possible to insert some css to adjust the appearance of the app, and override/translate the wording via a js configuration.
```yml
# cryptletter.yml
app:
  additional:
    css: './your/own/additional.css'
    js: './your/own/custom.js'
```

Further customization require a full build, since the assets are embedded into the executable for ease of use.

## Build

Run `./build.sh` and get your executable (you may adjust the docker build push destination).
