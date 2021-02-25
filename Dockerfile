# build
FROM golang:1.16-alpine3.12 as builder
WORKDIR /app
COPY . /app
ENV GO111MODULE=on \
    GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o geotool ./cmd/main.go
# runtime
FROM alpine:3.12
RUN apk add --no-cache tzdata
RUN apk add --no-cache ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /app/geotool .
ENTRYPOINT ["./geotool"]