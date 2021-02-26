FROM golang:1.15-alpine3.12 AS builder

COPY . /github.com/ivankonevv/pocket_telegram_bot/
WORKDIR /github.com/ivankonevv/pocket_telegram_bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/ivankonevv/pocket_telegram_bot/bin/bot .
COPY --from=0 /github.com/ivankonevv/pocket_telegram_bot/configs configs/

EXPOSE 80

CMD ["./bot"]
