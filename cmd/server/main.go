package main

import (
	"fmt"
	"homesync/server"
)

func main() {
	fmt.Println("Starting Homesync client and server")

	client := new(server.HomeSyncServer)
	client.Start()
}
