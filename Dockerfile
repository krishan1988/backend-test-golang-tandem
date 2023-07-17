FROM golang:1.19.2-alpine3.16

ADD config.json /tmp/btgConfig.json

ADD tools/public.key /tmp/public.key 

RUN mkdir /opt/backend-test-golang

ADD . /opt/backend-test-golang

WORKDIR /opt/backend-test-golang

RUN go build -o bin/backend-test-golang cmd/backend-test-golang/backend_test_golang.go