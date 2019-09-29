
FROM alpine:3.10

COPY ./entry_point_gitlab_token.sh /usr/src/myapp

WORKDIR "/usr/src/myapp"
ENTRYPOINT 'sh', 'entry_point_gitlab_token.sh'