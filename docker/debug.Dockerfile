FROM golang:alpine3.19

ENV TARGET="localhost"
ENV PORT_SCANS="nmap-tcp-top"
ENV SERVICE_SCANS="whatweb"

RUN CGO_ENABLED=0 go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

RUN apk add --no-cache git

RUN apk add --update-cache \
    nmap nmap-scripts \
    && rm -rf /var/cache/apk/*

COPY . .
COPY ./docker/scripts/debug.sh /debug.sh