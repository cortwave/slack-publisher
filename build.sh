#!/bin/bash

curl -o ca-certificates.crt https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt
docker build -t cortwave/slack-publisher:$1 .
rm ca-certificates.crt
rm main
