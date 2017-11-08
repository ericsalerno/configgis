# configgis

This is (will be) a super simple site+stage based Configuration server built in Go and using redis to get it's way.

## Usage

This will be replaced with a docker container and vars at some point. For now, enter the local port to listen on and the redis server connection info.

    package main

    import configgis "github.com/ericsalerno/configgis"

    func main() {
        server := configgis.NewServer(765, "localhost", 6379)
        server.Listen()
    }

## API

### Set Value

POST a value to the endpoint:

    /set/<sitename>/<stage>/<key>

### Get Value

Make a GET request to the endpoint:

    /get/<sitename>/<stage>/<key>

