FROM golang:1.19.10

RUN groupadd -g 1000 app
RUN useradd -g 1000 -u 1000 app
RUN mkdir -p /var/www /home/app/.cache

COPY ./ /var/www/.
WORKDIR /var/www

RUN go install github.com/go-delve/delve/cmd/dlv@master

RUN chown -R 1000:1000 /var/www /home/app /go /tmp
RUN chmod 777 /tmp
RUN chmod +t /tmp
USER app

EXPOSE 4000 8002
