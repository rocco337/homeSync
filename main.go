package main

import (
	"fmt"
	"homesync/client/homesyncserverservice"
)

func main() {
	fmt.Println("Starting Homesync client and server")

	// server := new(server.HomeSyncServer)
	// server.Start()

	serverService := new(homesyncserverservice.HomesyncServerService)
	serverService.GetFolderTree()

	// client := new(client.HomeSyncClient)
	// client.Start()

}
