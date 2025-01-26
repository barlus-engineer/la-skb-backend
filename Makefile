.PHONY: all clean build

all: build

build:
	go build -o out/laskb-server-api cmd/main.go
	cp config.yml out/

clean:
	rm -rf bin