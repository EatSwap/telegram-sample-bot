.ONESHELL:

build: sync-dep
	cd src
	go build -x -ldflags="-s -w" -o ../app telegram-sample-bot
	file ../app*

sync-dep:
	cd src
	go mod tidy

clean:
	rm app*
