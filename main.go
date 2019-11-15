package main

import (
	"fmt"
	"homesync/client"
)

func main() {
	fmt.Println("Starting Homesync client and server")

	// server := new(server.HomeSyncServer)
	// server.Start()

	client := new(client.HomeSyncClient)
	client.Start()
}
