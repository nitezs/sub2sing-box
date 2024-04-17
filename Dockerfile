FROM golang:1.21 as builder
LABEL authors="nite07"

WORKDIR /app

COPY . .
RUN go mod download

ARG version

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X sub2sing-box/constant.Version=${version}" -o sub2sing-box main.go

WORKDIR /app

FROM alpine:latest

COPY --from=builder /app/sub2sing-box /app/sub2sing-box
COPY --from=builder /app/templates /app/templates-origin
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

VOLUME [ "/app/templates" ]
EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]
