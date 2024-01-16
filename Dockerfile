FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-alpine-base:3.19-latest
#FROM ghcr.io/westdeutscherrundfunkkoeln/am-docker-oracle-ic-base:main-4
#COPY main app/main

COPY cmd/tim_web_show_apps/main app/main
COPY web app/web

WORKDIR /app

CMD ./main