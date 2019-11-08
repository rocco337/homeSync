package server

import (
	"fmt"
	"homesync/foldermonitor"
	"io"
	"os"
	"strings"
)

/* HardDriveOperations */
type HardDriveOperations struct {
	RootPath string
}

func (service HardDriveOperations) Create(info foldermonitor.FileInfo) {
	destinationPath := service.RootPath + "/" + info.RelativePath

	err := os.MkdirAll(strings.Replace(destinationPath, info.Name, "", 1), 0755)
	if err == nil || os.IsExist(err) {
	} else {
		panic(err)
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		panic(err)
	}

	source, err := os.Open(info.Path)
	if err != nil {
		return
	}
	defer source.Close()
	defer destination.Close()

	io.Copy(destination, source)
	fmt.Println("Soruce", info.Path, " is copied to ", destinationPath)
}

func (serivce HardDriveOperations) Remove(info foldermonitor.FileInfo) {

}

func (service HardDriveOperations) Tree() map[string]foldermonitor.FileInfo {
	monitorService := new(foldermonitor.FileMonitorService)
	monitorService.RootPath = service.RootPath
	return monitorService.Scan()
}
