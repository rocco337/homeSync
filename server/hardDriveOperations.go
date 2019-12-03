package server

import (
	"fmt"
	"homesync/foldermonitor"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const unixFolderPermission = 0777

/* HardDriveOperations */
type HardDriveOperations struct {
	RootPath string
}

/*Create - creates file on hard drive*/
func (service HardDriveOperations) Create(relativePath string, fileName string, fileStream multipart.File) {
	destinationPath := service.RootPath + "/" + relativePath

	//Creates
	err := os.MkdirAll(strings.Replace(destinationPath, fileName, "", 1), unixFolderPermission)
	if err == nil || os.IsExist(err) {
	} else {
		panic(err)
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		panic(err)
	}

	defer destination.Close()

	io.Copy(destination, fileStream)
	fmt.Println("Soruce", relativePath, " is copied to ", destinationPath)
}

/*Remove */
func (service HardDriveOperations) Remove(path string) {
	destinationPath := service.RootPath + "/" + path

	err := os.Remove(destinationPath)
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(destinationPath)
	if IsDirEmpty(dir) {
		err = os.Remove(dir)
		if err != nil {
			panic(err)
		}
	}
}

/*IsDirEmpty */
func IsDirEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	// read in ONLY one file
	n, err := f.Readdir(1)

	if len(n) <= 0 {
		return true
	}
	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true
	}
	return false
}

/*Tree - scans folder and return structure */
func (service HardDriveOperations) Tree(path string) map[string]foldermonitor.FileInfo {
	monitorService := new(foldermonitor.FileMonitorService)
	monitorService.RootPath = service.RootPath + "/" + path

	if _, err := os.Stat(monitorService.RootPath); os.IsNotExist(err) {
		err = os.Mkdir(monitorService.RootPath, unixFolderPermission)
		if err != nil {
			panic(err)
		}
	}

	return monitorService.Scan()
}
