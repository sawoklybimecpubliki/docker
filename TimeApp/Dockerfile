# Builder stage
FROM golang:latest AS builder

WORKDIR /app

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine

WORKDIR /bin

COPY --from=builder /app .

EXPOSE 8089

ENTRYPOINT ["/bin/app"]
