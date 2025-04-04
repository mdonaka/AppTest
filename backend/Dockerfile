FROM golang:1.24 AS dev

FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o main

FROM debian:stable-slim AS prod
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
