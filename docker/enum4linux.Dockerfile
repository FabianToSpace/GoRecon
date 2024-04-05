FROM golang:alpine3.19 as builder

RUN apk upgrade --update-cache --available \
    && apk git

RUN git clone https://github.com/cddmp/enum4linux-ng.git

FROM uptospace/gorecon:latest

WORKDIR /app

COPY --from=builder /enum4linux-ng /app

RUN apk upgrade --update-cache --available \
    && apk add samba-client python3 \
    && rm -rf /var/cache/apk/* \
    && ln -sf python3 /usr/bin/python

RUN python3 -m ensurepip

RUN pip3 install --no-cache-dir -r ./requirements.txt

RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gouser

ENTRYPOINT /go/bin/GoRecon ${TARGET}