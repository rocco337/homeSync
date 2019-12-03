package client

import (
	"fmt"
	"homesync/client/homesyncserverservice"
	"homesync/foldermonitor"
	"strconv"
	"time"
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
	serverService.Username = "rbobic"
	serverService.BaseUrl = "http://localhost:8080/"

	fmt.Println("Starting to monitor folder: " + LocalFolderPath)
	for {
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

		fmt.Println("UPLOAD: " + strconv.Itoa(len(filesToUpload)) + " files")
		for _, value := range filesToUpload {
			serverService.Upload(value)
		}

		fmt.Println("REMOVE: " + strconv.Itoa(len(filesToRemoveFromRemote)) + " files")
		for _, value := range filesToRemoveFromRemote {
			serverService.Remove(value)
		}

		time.Sleep(time.Second * 10)
	}

}

func throwAndLogIfError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		panic(err)
	}
}
