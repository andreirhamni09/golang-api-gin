package main

import (
	"api-gin/connections"
	"api-gin/handlers"
)

func main() {
	connections.Connection()
	handlers.HandlerReq()
}
