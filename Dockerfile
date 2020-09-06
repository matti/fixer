FROM golang:1.15.0-alpine3.12

RUN apk add --no-cache \
  bash

WORKDIR /build

COPY go.* ./
RUN go mod download

COPY . .
