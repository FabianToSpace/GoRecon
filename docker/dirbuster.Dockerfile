FROM golang:alpine3.19 as builder

RUN apk upgrade --update-cache --available \
    && apk add openssl

RUN wget https://github.com/epi052/feroxbuster/releases/latest/download/x86_64-linux-feroxbuster.zip -qO feroxbuster.zip \
    && unzip -d /tmp/ feroxbuster.zip feroxbuster \
    && chmod +x /tmp/feroxbuster

FROM uptospace/gorecon:latest

COPY --from=builder /tmp/feroxbuster /usr/local/bin/feroxbuster
RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gouser

ENTRYPOINT /go/bin/GoRecon ${TARGET}