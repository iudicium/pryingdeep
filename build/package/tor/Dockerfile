FROM alpine:latest

RUN apk update && apk add tor && mkdir -p /run/tor/service \
&& apk add curl && apk cache clean 

RUN chmod -R 700 /run/tor/service
WORKDIR /app


EXPOSE 9050

CMD ["tor"]

#  docker run     --rm       --name tor  --network host   --publish 9050:9050     tor