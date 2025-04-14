FROM golang:1.24.2-alpine3.21

# Setup server
WORKDIR /usr/src/app
COPY ./bin/cryptletter ./

EXPOSE 8080

CMD [ "./cryptletter", "serve"]
