FROM alpine:3.18

WORKDIR /app

COPY ./goapp /app/

ENTRYPOINT ./goapp