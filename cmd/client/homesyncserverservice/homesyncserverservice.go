package homesyncserverservice

import (
	"bytes"
	"encoding/json"
	"homesync/foldermonitor"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

//HomesyncServerService handle operations with remote server
type HomesyncServerService struct {
	BaseUrl  string
	Username string
}

//Upload files to remote server
func (service HomesyncServerService) Upload(info foldermonitor.FileInfo) {
	request, err := service.newfileUploadRequest(service.BaseUrl+"api/upload", info.Path, info.RelativePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		panic(err)
	}
}

//Remove file from server
func (service HomesyncServerService) Remove(info foldermonitor.FileInfo) {
	form := url.Values{}
	form.Add("pathToDelete", info.RelativePath)

	req, _ := http.NewRequest("POST", service.BaseUrl+"api/delete", strings.NewReader(form.Encode()))
	req.Header.Add("username", service.Username)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
}

/*GetFolderTree - calls remote server and gets state of remote folder */
func (service HomesyncServerService) GetFolderTree() map[string]foldermonitor.FileInfo {
	//should call remote server
	req, _ := http.NewRequest("GET", service.BaseUrl+"api/tree", nil)
	req.Header.Add("username", service.Username)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	//var result TreeResult
	var jsonResult map[string]map[string]foldermonitor.FileInfo
	err = json.Unmarshal(body, &jsonResult)

	return jsonResult["data"]
}

func (service HomesyncServerService) newfileUploadRequest(uri string, path string, filename string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("data", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	_ = writer.WriteField("relativePath", filename)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("username", service.Username)
	return req, err
}

type TreeResult struct {
	//data map[string]foldermonitor.FileInfo
	data   map[string]interface{} `json:"data"`
	status string                 `json:"status"`
}
