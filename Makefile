.PHONY:  #  – Указывает, что build и run не являются файлами, а просто командами для make
.SILENT:  # – Отключает вывод команд перед их выполнением (необязательно, можно удалить)

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot
