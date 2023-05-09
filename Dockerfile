FROM golang:1.19-alpine3.16 AS builder
#
WORKDIR /app
COPY . .

RUN GOOS=linux go build -o ./bin/whip-round ./cmd/main.go

# Run stage
# scratch is a special base image that is empty
FROM scratch:latest AS runner

COPY --from=builder /whip-round/bin/whip-round/ .
COPY --from=builder /whip-round/.env .
COPY --from=builder /whip-round/configs ./configs

EXPOSE 8080
CMD ["./whip-round"]