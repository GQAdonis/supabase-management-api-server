FROM golang:1.21-alpine3.19 AS base

RUN apk --update upgrade && apk --no-cache --update-cache --upgrade --latest add ca-certificates build-base gcc

WORKDIR /app

ADD go.mod go.mod
ADD go.sum go.sum

ENV GO111MODULE on
ENV CGO_ENABLED 1

RUN go mod download
RUN go mod tidy

ADD . .

RUN go build -o /usr/bin/server cmd/main.go

FROM alpine:3.19

COPY --from=base /usr/bin/server /usr/bin/server

EXPOSE 8180

ENTRYPOINT ["/usr/bin/server"]

