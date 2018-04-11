FROM golang:1.9.3-alpine
MAINTAINER Kiran Pachhai <kiran.pachhai@gmail.com>

ENV server_path /go/src/gitlab.com/kpachhai/bytecubed-challenge/
ENV PATH $PATH:$server_path

WORKDIR $server_path
COPY . .

RUN go build -o webserver .

ENTRYPOINT ./webserver

EXPOSE 80 443


