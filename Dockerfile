FROM golang:alpine
MAINTAINER Piotr Łuczak<piotrluczak1995@gmail.com>

COPY screens2tex.go .
RUN go build screens2tex.go

ENTRYPOINT ["./screens2tex"]
