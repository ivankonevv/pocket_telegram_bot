.PHONY:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t pocket_bot:v1.0 .

start-container:
	docker run --name pocket_bot -p 80:80 --env-file .env pocket_bot:v1.0