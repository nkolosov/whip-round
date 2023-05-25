FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux go build -o ./bin/whip-round ./cmd/main.go

FROM alpine:latest AS runner

WORKDIR /app
COPY --from=builder /app/bin/whip-round .
COPY --from=builder /app/schema ./schema
COPY --from=builder /app/.env ./.env

EXPOSE 8080
CMD ["./whip-round"]