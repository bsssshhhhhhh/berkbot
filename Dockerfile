FROM golang:1.16-alpine

RUN apk update
RUN apk upgrade
RUN apk add bash

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN /bin/bash build.sh

CMD ["/app/build/berkbot"]