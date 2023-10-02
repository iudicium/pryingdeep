FROM alpine:latest

RUN apk update && apk add tor

COPY torrc /etc/tor/torrc 

RUN mkdir -p /run/tor && chown tor /run/tor && chmod 0700 /run/tor


USER tor



ENTRYPOINT ["tor"]


