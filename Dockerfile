FROM alpine:3.4

EXPOSE 8080

ARG app=go-linux-amd64

COPY $app /app

ENTRYPOINT ["/app"]
