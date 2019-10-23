
FROM alpine:3.10

RUN apk add --no-cache git \
   && rm -rf /var/cache/* \
   && rm -rf /root/.cache/*

COPY ./entry_point_gitlab_token.sh /usr/src/myapp

WORKDIR "/usr/src/myapp"
ENTRYPOINT 'sh', 'entry_point_gitlab_token.sh'