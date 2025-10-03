FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o my-go-api main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root
COPY --from=builder /app/my-go-api .
COPY .env .
EXPOSE 8080
CMD ["./my-go-api"]