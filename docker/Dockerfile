FROM golang:1.15.2-alpine AS base

ENV GOPATH '/'
WORKDIR /project

COPY ./api .
RUN apk add --no-cache git \
 && go build -o api

FROM alpine AS main

WORKDIR /project/api
COPY --from=base /project/api .

CMD ["./api"]
