package main

import (
	"fmt"

	"homesync.com/foldermonitor"
)

func main() {
	localFolderPath := "/home/roko/sharedTest"
	remoteFolderPath := "/home/roko/sharedTestRemote"

	localFileMonitorService := new(foldermonitor.FileMonitorService)
	localFileMonitorService.RootPath = localFolderPath

	remoteFileMonitorService := new(foldermonitor.FileMonitorService)
	remoteFileMonitorService.RootPath = remoteFolderPath

	fmt.Println("Starting to monitor folder: " + localFolderPath)
	localFiles := localFileMonitorService.Scan()
	remoteFiles := remoteFileMonitorService.Scan()

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
	for key, value := range filesToUpload {
		fmt.Println(key, value.Path, value.Modified)
	}

	fmt.Println("REMOVE=====================")
	for key, value := range filesToRemoveFromRemote {
		fmt.Println(key, value.Path, value.Modified)
	}
}

func throwAndLogIfError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		panic(err)
	}
}

type FileServerService struct {
	rootPath string
}

func (serivce FileServerService) Add(file []byte, info foldermonitor.FileInfo) {

}

func (serivce FileServerService) Remove(file []byte, info foldermonitor.FileInfo) {

}
