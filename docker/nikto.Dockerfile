FROM uptospace/gorecon:0.0.6

RUN apk upgrade --update-cache --available \
    && apk add nikto \
    && rm -rf /var/cache/apk/*

RUN cp /usr/bin/nikto.pl /usr/local/bin/nikto

RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gorecon

ENTRYPOINT /go/bin/GoRecon ${TARGET}