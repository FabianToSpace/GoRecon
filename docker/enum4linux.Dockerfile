FROM alpine:3.19 as builder

RUN apk upgrade --update-cache --available \
    && apk add git

RUN git clone https://github.com/cddmp/enum4linux-ng.git
RUN pwd

FROM uptospace/gorecon:0.0.6

RUN apk upgrade --update-cache --available \
    && apk add samba-client python3 py3-samba py3-ldap3 py3-yaml py3-impacket \
    && rm -rf /var/cache/apk/* \
    && ln -sf python3 /usr/bin/python

COPY --from=builder /enum4linux-ng/enum4linux-ng.py /usr/local/bin/enum4linux-ng

RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gorecon

ENTRYPOINT /go/bin/GoRecon ${TARGET}