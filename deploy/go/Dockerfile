FROM golang:1.21.6-alpine3.19 as builder

# install alpine packages
RUN apk update \
    && apk add make

# copy project files
COPY ./src/api/. /app
WORKDIR /app

# build app
RUN go build ./cmd/app

# run migrations and load fixtures
RUN make migrate-up \
    && make fixtures-load


FROM scratch
COPY --from=builder /app/.env /app/
COPY --from=builder /app/app /app/
CMD ["/app/app"]
