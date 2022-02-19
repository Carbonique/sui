
FROM busybox

WORKDIR /app/dashboard
COPY . /app/dashboard

EXPOSE 80

ENTRYPOINT [ "httpd", "-f", "-v", "-h", "/app/dashboard", "-u", "1000" ]
