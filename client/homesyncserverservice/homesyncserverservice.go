package homesyncserverservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homesync/foldermonitor"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type HomesyncServerService struct {
}

func (service HomesyncServerService) Upload(info foldermonitor.FileInfo) {
	request, err := newfileUploadRequest("http://localhost:8080/api/upload", info.Path, info.RelativePath)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	fmt.Println(resp)
}

func newfileUploadRequest(uri string, path string, filename string) (*http.Request, error) {
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
	return req, err
}

func (serivce HomesyncServerService) Remove(info foldermonitor.FileInfo) {
	response, err := http.PostForm("http://localhost:8080/api/delete", url.Values{
		"pathToDelete": {info.RelativePath},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

/*GetFolderTree - calls remote server and gets state of remote folder */
func (service HomesyncServerService) GetFolderTree() map[string]foldermonitor.FileInfo {
	//should call remote server
	resp, err := http.Get("http://localhost:8080/api/tree")
	if err != nil {
		// handle error
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

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
