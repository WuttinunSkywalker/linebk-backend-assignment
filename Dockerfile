# Build stage
FROM golang:1.24.4-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
EXPOSE 8080
CMD [ "/app/main" ]