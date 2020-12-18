FROM golang:1.15.0-alpine3.12

RUN apk add --no-cache \
  git

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY . .
