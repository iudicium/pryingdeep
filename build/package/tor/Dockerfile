FROM alpine:latest

RUN apk update && apk add tor && mkdir -p /run/tor/service \
&& apk add curl && apk cache clean 
ADD ./entrypoint.sh /

RUN chmod +x /entrypoint.sh

WORKDIR /app


EXPOSE 9050

ENTRYPOINT ["/entrypoint.sh"]


#  docker run     --rm       --name tor  --network host   --publish 9050:9050     app