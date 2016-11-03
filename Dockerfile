FROM alpine:3.4

EXPOSE 8080

ARG binname=go-linux-amd64

COPY $binname /app

ENTRYPOINT ["/app"]
