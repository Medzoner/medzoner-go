#############################################
### Base build
#############################################
FROM golang:1.21.9 AS base

RUN groupadd -g 1000 app && \
    useradd -g 1000 -u 1000 app

COPY ./ /var/www/.

RUN go install github.com/go-delve/delve/cmd/dlv@master

#############################################
### Run the Go Binary
#############################################
FROM golang:1.21.9

COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /go/bin/dlv /

RUN mkdir -p /var/www && \
    mkdir -p /home/app/.cache

COPY --from=base /var/www /var/www
RUN chown -R 1000:1000 /var/www /home/app /tmp
RUN chmod 777 /tmp
RUN chmod +t /tmp

USER app

WORKDIR /var/www

EXPOSE 4000 8002
