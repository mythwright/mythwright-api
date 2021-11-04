# Mythwright - API
The API backend for the Mythwright frontend UI

## Building
```shell
$ go build -o mythwright cmd/mythwright-api/main.go
```

## Architecture
- Two OAuth Providers: Discord & GW2Auth
- MongoDB database storage
- ReactJS frontend
- Go based Relay application to handle private DBs

## Features
* Create custom item mappings, map bags to their drops (or a subset of them)
* Add Datasets from Salvaging, Farming, Crafting, etc.
* Login with Discord for easy use, or basic auth to avoid OAuth
* Keep Datasets private utilizing a Mythwright Relay (all data is stored on your private db, requests are proxied to the self-hosted relay)
