FROM alpine:3.21.2

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /migrations

ADD *.sql /migrations/
ADD entrypoint.sh /migrations/

RUN ["chmod", "+x", "/migrations/entrypoint.sh"]

ENTRYPOINT ["/migrations/entrypoint.sh"]
