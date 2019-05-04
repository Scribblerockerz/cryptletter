FROM golang:alpine

# Setup server
WORKDIR /usr/src/app
COPY ./bin/cryptletter ./
COPY ./theme ./theme
COPY ./public ./public

EXPOSE 8080

CMD [ "./cryptletter"]
