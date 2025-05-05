# Builder stage
FROM golang:latest AS builder

WORKDIR /app1

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine

WORKDIR /bin

COPY --from=builder /app1 .

EXPOSE 8088

ENTRYPOINT ["/bin/app"]
