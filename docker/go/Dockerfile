FROM golang:1.21.1-alpine3.18 as builder

# install alpine packages
RUN apk update \
    && apk add curl \
    && apk add make

# install golang-migrate/migrate
RUN cd /home \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz \
    && cp /home/migrate /bin

# copy project files
COPY ./src/api/. /app
WORKDIR /app

# build app
RUN go build ./cmd/app

# run migrations
RUN make migrate-up


FROM scratch
COPY --from=builder /app/.env /app/
COPY --from=builder /app/app /app/
CMD ["/app/app"]
