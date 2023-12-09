FROM alpine:latest

RUN apk update && \
    apk add tor && \
    mkdir -p /run/tor/service && \
    chmod -R 700 /run/tor/service && \
    apk add curl && \
    apk cache clean

WORKDIR /app

EXPOSE 9050

CMD ["tor"]
