FROM golang:alpine3.19

ENV TARGET="localhost"
ENV PORT_SCANS="nmap-tcp-top"
ENV SERVICE_SCANS="whatweb"

RUN go install github.com/FabianToSpace/GoRecon@v0.0.5

RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

RUN ln -s /go/bin/GoRecon /usr/local/bin/gorecon \
    && ln -s /go/bin/GoRecon /usr/local/bin/GoRecon

USER gorecon

ENTRYPOINT /go/bin/GoRecon ${TARGET}