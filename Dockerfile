FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go build -o dns-resolver

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/dns-resolver .
CMD ["./dns-resolver"]
