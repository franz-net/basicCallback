FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum
#RUN go mod download
COPY . .
RUN go build -o callback-server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/callback-server .

EXPOSE 8080

CMD ["./callback-server"]
