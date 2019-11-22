package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Homesync client and server")

	// server := new(server.HomeSyncServer)
	// server.Start()

	client := new(HomeSyncClient)
	client.Start()
}
