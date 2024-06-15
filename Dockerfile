FROM alpine:3.19

LABEL org.opencontainers.image.source https://github.com/carbonique/sui

RUN apk add --no-cache busybox-extras

WORKDIR /app/dashboard
COPY ./web /app/dashboard/

EXPOSE 80

ENTRYPOINT [ "httpd", "-f", "-v", "-h", "/app/dashboard", "-u", "1000"]
