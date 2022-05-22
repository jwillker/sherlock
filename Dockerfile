FROM golang:1.18-alpine as builder

RUN apk update && \
    apk add --no-cache git && \
    rm -rf /var/cache/apk/*

WORKDIR /app/
COPY . /app/

RUN go mod download && \
    go build -o sherlock


FROM alpine

COPY --from=builder /app/sherlock /sherlock

ENTRYPOINT [ "/sherlock" ]
