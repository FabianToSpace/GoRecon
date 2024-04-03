FROM golang:alpine3.19

ENV TARGET localhost
ENV PORT_SCANS nmap-tcp-top
ENV SERVICE_SCANS whatweb

RUN apk add --no-cache git

RUN go install github.com/FabianToSpace/GoRecon@latest

ENTRYPOINT ["/go/bin/GoRecon", "${TARGET}"]