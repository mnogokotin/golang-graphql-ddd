FROM golang:1.21.1-alpine3.18

# create the appropriate directories
RUN mkdir -p /home/app
ENV APP_HOME=/home/app/web
RUN mkdir $APP_HOME
WORKDIR $APP_HOME

# install gcc for hiredis
RUN apk update \
    && apk add curl \
    && apk add make
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz \
    && cp migrate /bin/

# copy project
COPY ./docker/entrypoint.sh $APP_HOME
COPY ./src/api/. $APP_HOME

# build app
RUN go build ./cmd/app/

# run entrypoint
ENTRYPOINT ["/home/app/web/entrypoint.sh"]
