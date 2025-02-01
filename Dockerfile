FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/main .

RUN ls -l /app

RUN chmod +x /app/main

EXPOSE 8080

CMD ["/app/main"]
