FROM golang:latest

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy