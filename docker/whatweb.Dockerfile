FROM uptospace/gorecon:0.0.6 as gorecon

FROM ruby:2.7-alpine as builder

RUN apk upgrade --update-cache --available \
    && apk add git make gcc musl-dev \
    && rm -rf /var/cache/apk/*

RUN git clone https://github.com/urbanadventurer/WhatWeb.git /whatweb

RUN gem install rchardet:1.8.0 mongo:2.14.0 json:2.5.1

RUN cd /whatweb && bundle update && bundle install

FROM ruby:2.7-alpine

RUN apk upgrade --update-cache --available \
    && apk add git make \
    && rm -rf /var/cache/apk/*

COPY --from=builder /usr/local/bundle/ /usr/local/bundle/
COPY --from=builder /whatweb /whatweb/

COPY --from=gorecon /go /go

RUN ln -s /whatweb/whatweb /usr/local/bin/whatweb


# RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

# USER gorecon

ENTRYPOINT /go/bin/GoRecon ${TARGET}