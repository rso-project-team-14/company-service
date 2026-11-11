
# Build the Go application
FROM golang:1.27.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /company-service .

# Run the compiled application
FROM alpine:3.24

RUN apk add --no-cache ca-certificates

COPY --from=builder /company-service /company-service

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/company-service"]
