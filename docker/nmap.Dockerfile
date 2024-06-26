FROM uptospace/gorecon:latest

RUN apk upgrade --update-cache --available \
    && apk add nmap nmap-scripts \
    && rm -rf /var/cache/apk/*

RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gouser

ENTRYPOINT /go/bin/GoRecon ${TARGET}