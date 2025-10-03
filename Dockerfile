FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o canvas-api main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root
COPY --from=builder /app/canvas-api .
COPY .env .
EXPOSE 4000
CMD ["./canvas-api"]