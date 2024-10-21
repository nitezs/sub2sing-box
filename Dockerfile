FROM golang:1.23 as builder
LABEL authors="nite07"

WORKDIR /app
COPY . .
RUN go mod download
ARG version
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X github.com/nitezs/sub2sing-box/constant.Version=${version}" -o sub2sing-box .
WORKDIR /app

FROM alpine:latest
COPY --from=builder /app/sub2sing-box /app/sub2sing-box
VOLUME [ "/app/templates" ]
EXPOSE 8080
CMD ["/app/sub2sing-box", "server"]
