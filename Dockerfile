FROM golang:1.21 as builder
LABEL authors="nite07"

WORKDIR /app

COPY . .
RUN go mod download

ARG version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X sub2clash/config.Version=${version}" -o sub2sing-box main.go

FROM alpine:latest

COPY --from=builder /app/sub2sing-box /app/sub2sing-box
COPY --from=builder /app/template /app/template

ENTRYPOINT ["/app/sub2sing-box","server"]
