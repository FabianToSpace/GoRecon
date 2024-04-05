FROM epi052/feroxbuster as ferox

FROM uptospace/gorecon:0.0.6

COPY --from=ferox /usr/local/bin/feroxbuster /usr/local/bin/feroxbuster
RUN adduser -D gorecon && chown -R gorecon:gorecon /go/bin 

USER gorecon

ENTRYPOINT /go/bin/GoRecon ${TARGET}