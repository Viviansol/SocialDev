FROM golang:1.18.1-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN go build -buildvcs=false -o bin/fresh_market

FROM alpine:3.16

RUN apk update

ENV DOCKERIZE_VERSION v0.6.1

RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

EXPOSE 8080

COPY --from=builder /usr/src/app/bin .

COPY ./.env .