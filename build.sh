#!/bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
curl -o ca-certificates.crt https://curl.haxx.se/ca/cacert.pem
docker build -t cortwave/slack-publisher:$1 .
rm ca-certificates.crt
rm main
