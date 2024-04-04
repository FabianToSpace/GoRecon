FROM golang:alpine3.19

ENV TARGET="localhost"
ENV PORT_SCANS="nmap-tcp-top"
ENV SERVICE_SCANS="whatweb"

RUN apk add --no-cache git

RUN apk add --update-cache \
    nmap nmap-scripts \
    && rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .

RUN go build -o GoRecon && chmod +x GoRecon && mv GoRecon /go/bin

RUN adduser -D gorecon && chown -R gorecon:gorecon /app /go/bin 
USER gouser

ENTRYPOINT /go/bin/GoRecon ${TARGET}