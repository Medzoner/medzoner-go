FROM alpine AS runner

COPY ./bin /opt/bin
COPY ./tmpl /opt/tmpl
COPY ./var /opt/var
COPY ./public /opt/public
COPY ./migrations /opt/migrations
COPY ./pkg/infra/resources/data /opt/pkg/infra/resources/data

RUN addgroup --gid 1000 -S app && adduser --uid 1000 -S -G app app

RUN chown -R app:app /opt

USER app

WORKDIR /opt

CMD ["/opt/bin/app"]
