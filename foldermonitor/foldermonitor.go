package foldermonitor

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileMonitorService struct {
	lastTimeChecked time.Time
	RootPath        string
}

func (service FileMonitorService) Scan() map[string]FileInfo {
	var files = make(map[string]FileInfo)

	err := filepath.Walk(service.RootPath, func(path string, info os.FileInfo, err error) error {
		throwAndLogIfError(err)
		if !info.IsDir() {
			file := new(FileInfo)
			file.RelativePath = strings.Replace(path, service.RootPath, "", 1)
			file.Path = path
			file.Modified = info.ModTime()
			file.Name = info.Name()

			hashKey := file.GetContentHash()
			files[hashKey] = *file
		}
		return nil
	})

	throwAndLogIfError(err)

	return files
}

type FileInfo struct {
	Path         string
	RelativePath string
	Modified     time.Time
	Name         string
}

func (info FileInfo) GetContentHash() string {
	file, err := os.Open(info.Path)
	throwAndLogIfError(err)

	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		throwAndLogIfError(err)
	}

	hash.Write([]byte(info.RelativePath))
	hashInBytes := hash.Sum(nil)[:16]

	return hex.EncodeToString(hashInBytes)
}

func throwAndLogIfError(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		panic(err)
	}
}
