FROM golang:alpine

# Setup server
WORKDIR /usr/src/app
COPY ./bin/cryptletter ./

EXPOSE 8080

CMD [ "./cryptletter", "serve"]
