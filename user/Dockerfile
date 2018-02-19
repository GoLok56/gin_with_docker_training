FROM golang:1.8

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/gin-gonic/gin
RUN go get github.com/spf13/cast

ADD . /go/src/github.com/golok56/gin_with_docker_training/user

WORKDIR /go/src/github.com/golok56/gin_with_docker_training/user

RUN go install github.com/golok56/gin_with_docker_training/user

ENTRYPOINT /go/bin/user

EXPOSE 8080