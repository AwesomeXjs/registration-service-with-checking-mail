FROM golang:1.18-alpine AS builder

RUN apk add --no-cache bash git

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

WORKDIR /root

COPY internal/migrations/*.sql migrations/
COPY scripts/migration.sh .
COPY .env .

RUN chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]