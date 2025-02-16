FROM golang:1.23.2-alpine as builder

WORKDIR /tasks-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copiar o arquivo .env
COPY .env .env

RUN go install github.com/jackc/tern/v2@latest

COPY ./migrate.sh /usr/local/bin/migrate.sh

RUN chmod +x /usr/local/bin/migrate.sh

RUN go build -o app ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /tasks-api/app /usr/local/bin/app
COPY --from=builder /usr/local/bin/migrate.sh /usr/local/bin/migrate.sh

COPY --from=builder /go/bin/tern /usr/local/bin/tern

# Copiar o .env para o contêiner final
COPY --from=builder /tasks-api/config/migrations /tasks-api/migrations
COPY --from=builder /tasks-api/.env /tasks-api/.env

WORKDIR /tasks-api

CMD /usr/local/bin/migrate.sh && /usr/local/bin/app
