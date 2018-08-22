#! /bin/bash

# 编译web、api、stream、dispatcher四个服务

cd ${GOPATH}/src/daker.wang/Azen/Go-execise/Streaming/api
env GOOS=linux GOARCH=amd64 go build -o ../../bin/api

cd ${GOPATH}/src/daker.wang/Azen/Go-execise/Streaming/dispatcher
env GOOS=linux GOARCH=amd64 go build -o ../../bin/dispatcher

cd ${GOPATH}/src/daker.wang/Azen/Go-execise/Streaming/stream
env GOOS=linux GOARCH=amd64 go build -o ../../bin/stream

cd ${GOPATH}/src/daker.wang/Azen/Go-execise/Streaming/web
env GOOS=linux GOARCH=amd64 go build -o ../../bin/web