
.PHONY: build clean windows

build:
	GOOS=linux GOARCH=amd64 go build -o mythwright-api cmd/mythwright-api/main.go

windows:
	$$env:GOOS='linux'; $$env:GOARCH='amd64'; go build -o mythwright-api cmd/mythwright-api/main.go

clean:
	rm -f mythwright-api