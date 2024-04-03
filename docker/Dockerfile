FROM golang:alpine3.19

RUN apk add --no-cache git

RUN go install https://github.com/FabianToSpace/GoRecon@latest

ENTRYPOINT ["/go/bin/GoRecon"]