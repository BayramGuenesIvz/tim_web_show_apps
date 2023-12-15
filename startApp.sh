#!/bin/bash

rm ./cmd/tim_web_show_apps/main
cd ./cmd/tim_web_show_apps
CGO_ENABLED=0 GOOS=linux go build -o main
cd ..
cd ..
./cmd/tim_web_show_apps/main confLocation=./config/configshowapps.json