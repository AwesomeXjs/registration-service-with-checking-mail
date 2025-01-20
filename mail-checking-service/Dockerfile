FROM golang:1.23 AS builder

WORKDIR /usr/local/src

RUN apt-get update && apt-get install -y gcc libc6-dev

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

RUN go build -o ./bin/app ./cmd/mail/main.go

# Используем тот же базовый образ для runtime
FROM golang:1.23 AS runtime

WORKDIR /

COPY --from=builder /usr/local/src/bin/app .
COPY .env .

CMD ["./app"]
