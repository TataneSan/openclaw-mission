# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-api ./cmd/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/auth-api .

EXPOSE 8080

ENV PORT=8080

CMD ["./auth-api"]
