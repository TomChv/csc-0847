FROM golang:1.20-buster AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o api main.go

FROM debian:buster-slim AS APP

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/api ./

ENTRYPOINT ["./api"]