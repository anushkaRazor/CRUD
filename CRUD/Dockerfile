FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o todo ./cmd

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/todo .

RUN mkdir -p /app/log

EXPOSE 8080
CMD ["./todo"]
