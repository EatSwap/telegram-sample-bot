.ONESHELL:

build: sync-dep
	cd src
	go build -x -ldflags="-s -w" -o ../app telegram-sample-bot

sync-dep:
	cd src
	go mod tidy

clean:
	rm app*
