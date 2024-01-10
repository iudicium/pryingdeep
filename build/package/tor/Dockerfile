FROM alpine:3.19.0


RUN addgroup -S appgroup && adduser -S pryingdeep -G appgroup

WORKDIR /home/pryingdeep


RUN apk update && \
    apk add tor && \
    mkdir -p /run/tor/service && \
    chmod -R 700 /run/tor/service && \
    apk add curl && \
    apk cache clean

USER pryingdeep

WORKDIR /app

EXPOSE 9050

CMD ["tor"]
