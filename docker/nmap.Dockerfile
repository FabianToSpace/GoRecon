FROM uptospace/gorecon:0.0.6

RUN apk upgrade --update-cache --available \
    && apk add nmap nmap-scripts \
    && rm -rf /var/cache/apk/*

RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gorecon

ENTRYPOINT /go/bin/GoRecon ${TARGET}