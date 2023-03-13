FROM golang:1.20-buster AS builder

# Install ImageMagick dev library
RUN apt-get update && apt-get -q -y install libmagickwand-dev

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o thumbnails main.go

FROM debian:buster-slim AS APP

# Install ImageMagick deps and ca-certificates
RUN apt-get update && apt-get -q -y install \
	ca-certificates \
	libmagickwand-6.q16-6 && \
	rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/thumbnails ./

ENTRYPOINT ["./thumbnails"]