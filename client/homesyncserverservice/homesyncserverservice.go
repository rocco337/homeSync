package homesyncserverservice

import (
	"encoding/json"
	"fmt"
	"homesync/foldermonitor"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type HomesyncServerService struct {
	RootPath string
}

func (service HomesyncServerService) Upload(info foldermonitor.FileInfo) {
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

func (serivce HomesyncServerService) Remove(info foldermonitor.FileInfo) {

}

/*GetFolderTree - calls remote server and gets state of remote folder */
func (service HomesyncServerService) GetFolderTree() map[string]foldermonitor.FileInfo {
	//should call remote server
	resp, err := http.Get("http://localhost:8080/api/tree")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//var result TreeResult
	var jsonResult map[string]map[string]foldermonitor.FileInfo
	err = json.Unmarshal(body, &jsonResult)

	fmt.Println(jsonResult["data"])
	return jsonResult["data"]
}

type TreeResult struct {
	//data map[string]foldermonitor.FileInfo
	data   map[string]interface{} `json:"data"`
	status string                 `json:"status"`
}
