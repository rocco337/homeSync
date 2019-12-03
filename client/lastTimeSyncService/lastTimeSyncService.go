package lastTimeSyncService

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type LastTimeSyncService struct {
	ConfigFolderPath     string
	LastSyncTimeFilename string
}

//CreateConfigFileIfNotExist ...
func (service LastTimeSyncService) CreateConfigFileIfNotExist() {
	fmt.Println(service.ConfigFolderPath)
	if _, err := os.Stat(service.ConfigFolderPath); os.IsNotExist(err) {
		err := os.MkdirAll(service.ConfigFolderPath, 0666)
		if err != nil {
			panic("Cannot create config folder: " + service.ConfigFolderPath)
		}
	}
}

//CreateFileIfNotExist ...
func (service LastTimeSyncService) CreateFileIfNotExist(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		_, err := os.Create(filePath)
		if err != nil {
			panic("Cannot create file: " + service.ConfigFolderPath)
		}
	}
}

//Get ...
func (service LastTimeSyncService) Get() time.Time {

	service.CreateConfigFileIfNotExist()
	service.CreateFileIfNotExist(service.LastSyncTimeFilename)

	file, err := os.OpenFile(service.LastSyncTimeFilename, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		panic("Cannot read " + service.LastSyncTimeFilename)
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)
	fileContentString := string(buf.Bytes())

	if len(fileContentString) > 0 {
		parsedTime, err := time.Parse(time.RFC3339, fileContentString)
		if err != nil {
			panic("Cannot parse time  " + fileContentString)
		}
		return parsedTime
	}

	return time.Time{}
}

//Set ...
func (service LastTimeSyncService) Set(lastSyncTime time.Time) {
	path := service.ConfigFolderPath + "/" + service.LastSyncTimeFilename

	ioutil.WriteFile(path, []byte(lastSyncTime.Format(time.RFC3339)), 0666)
}
