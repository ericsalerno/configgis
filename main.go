package main

import (
	"os"
	"strconv"
)

func main() {

	localPortString := os.Getenv("CFIGGIS_PORT")
	redisHostString := os.Getenv("CFIGGIS_REDIS_HOST")
	redisPortString := os.Getenv("CFIGGIS_REDIS_PORT")

	localPort := 765
	redisPort := 6379

	if localPortString != "" {
		localPort, _ = strconv.Atoi(localPortString)
	}

	if redisHostString == "" {
		redisHostString = "localhost"
	}

	if redisPortString != "" {
		redisPort, _ = strconv.Atoi(redisPortString)
	}

	server := NewServer(localPort, redisHostString, redisPort)
	server.Listen()
}
