FROM golang:1.22-alpine3.19 as go-builder

RUN apk add build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG GOARCH=amd64
ARG GOOS=linux
ARG VERSION=0.1.0

ENV GOARCH=$GOARCH
ENV GOOS=$GOOS
ENV VERSION=$VERSION
ENV CGO_ENABLED=0
RUN go build -ldflags "-s -w -X main.Version=$VERSION -X main.Build=$(date +%s)" -o ipcameras-api

FROM --platform=linux/amd64 alpine:3.19
LABEL org.opencontainers.image.source "https://github.com/aicacia/ipcameras"

WORKDIR /opt/app

ARG APP_ENV=prod
ENV APP_ENV=${APP_ENV}

COPY --from=go-builder /app/ipcameras-api /opt/app/ipcameras-api

CMD [ "sh", "-c", "/opt/app/ipcameras-api" ]