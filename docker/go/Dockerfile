FROM golang:alpine

COPY ./src /go/src

WORKDIR /go/src/

RUN apk update \
  && apk add --no-cache git \
  && go get -v -u  github.com/gin-gonic/gin \
                   github.com/jinzhu/gorm \
                   github.com/go-sql-driver/mysql \
                   github.com/joho/godotenv \
                   github.com/google/uuid

EXPOSE 8080
