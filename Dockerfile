FROM golang:1.18-alpine as builder

RUN apk update && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/*

RUN go mod download && \
    go build -o /app/sherlock


FROM alpine

COPY --from=builder /app/sherlock /sherlock

ENTRYPOINT [ "sherlock" ]
