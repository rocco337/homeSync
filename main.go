package main

import (
	"fmt"
	"homesync/client"
)

func main() {
	fmt.Println("Starting Homesync client and server")

	//server := new(server.HomeSyncServer)
	//server.Start()

	// serverService := new(homesyncserverservice.HomesyncServerService)

	// file := new(foldermonitor.FileInfo)
	// file.Path = "/home/roko/Documents/Roko Bobic_CV.pdf"
	// file.RelativePath = "Roko Bobic_CV.pdf"
	// serverService.Upload(*file)

	client := new(client.HomeSyncClient)
	client.Start()

}
