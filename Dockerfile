FROM golang:1.22.6

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy