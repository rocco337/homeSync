package client

import (
	"fmt"
	"homesync/client/homesyncserverservice"
	"homesync/foldermonitor"
)

const LocalFolderPath = "/home/roko/sharedTest"

/*HomeSyncClient */
type HomeSyncClient struct {
}

/*Start - starts to monitor changes in root folder*/
func (client HomeSyncClient) Start() {

	localFileMonitorService := new(foldermonitor.FileMonitorService)
	localFileMonitorService.RootPath = LocalFolderPath

	//create server object - upload changed files
	serverService := new(homesyncserverservice.HomesyncServerService)

	fmt.Println("Starting to monitor folder: " + LocalFolderPath)
	localFiles := localFileMonitorService.Scan()
	remoteFiles := serverService.GetFolderTree()

	filesToUpload := make(map[string]foldermonitor.FileInfo)
	filesToRemoveFromRemote := make(map[string]foldermonitor.FileInfo)

	//find files to upload
	for key, value := range localFiles {
		if _, exists := remoteFiles[key]; !exists {
			filesToUpload[key] = value
		}
	}

	//find files to delete
	for key, value := range remoteFiles {
		if _, exists := localFiles[key]; !exists {
			filesToRemoveFromRemote[key] = value
		}
	}

	//find removed files
	fmt.Println("LOCAL=====================")
	for key, value := range localFiles {
		fmt.Println(key, value.RelativePath, value.Modified)
	}

	fmt.Println("UPLOAD=====================")
	for _, value := range filesToUpload {
		serverService.Upload(value)
	}

	fmt.Println("REMOVE=====================")
	for key, value := range filesToRemoveFromRemote {
		fmt.Println(key, value.Path, value.Modified)
		serverService.Remove(value)
	}
}

func throwAndLogIfError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		panic(err)
	}
}
