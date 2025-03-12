FROM golang:1.24 AS builder

WORKDIR /usr/src/maindir

COPY go,mod ./
RUN go mod download

COPY main.go .
RUN go build -v -o /usr/local/bin/gotka ./...

CMD [ "gotka" ]
