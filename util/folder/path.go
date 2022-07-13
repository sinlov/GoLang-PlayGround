package folder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetCurrentExecPath get exec path
func GetCurrentExecPath() (string, error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(s, "\\")
	p := string(s[0 : i+1])
	return p, nil
}

// GetCurrentFilePath run path
func GetCurrentFilePath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return file, nil
}

// PathExists path exists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// PathExistsFast path exists fast
func PathExistsFast(path string) bool {
	exists, _ := PathExists(path)
	return exists
}

// PathJoin is path.Join()
func PathJoin(elem ...string) string {
	return path.Join(elem...)
}

// PathParent as filepath.Dir
func PathParent(path string) string {
	return filepath.Dir(path)
}

// PathFolderList list of Path Folder
func PathFolderList(path string) []string {
	exists, err := PathExists(path)
	if err != nil || !exists {
		return nil
	}

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	if len(fileInfos) == 0 {
		return nil
	}
	var dirList []string
	for _, info := range fileInfos {
		if info.IsDir() {
			dirList = append(dirList, info.Name())
		}
	}
	return dirList
}

// PathFileList list of Path file
func PathFileList(path string) []string {
	exists, err := PathExists(path)
	if err != nil || !exists {
		return nil
	}

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	if len(fileInfos) == 0 {
		return nil
	}
	var dirList []string
	for _, info := range fileInfos {
		if !info.IsDir() {
			dirList = append(dirList, info.Name())
		}
	}
	return dirList
}

// ReadFileAsJson read file as json
func ReadFileAsJson(path string, v interface{}) error {
	exists, err := PathExists(path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("path not exist %v", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

// WriteFileAsJson write json file
// path most use Abs Path
// v data
// fileMod os.FileMode(0666) or os.FileMode(0644)
// coverage true will coverage old
// beauty will format json when write
func WriteFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if beauty {
		var str bytes.Buffer
		err := json.Indent(&str, data, "", "  ")
		if err != nil {
			return err
		}
		err = os.WriteFile(path, str.Bytes(), fileMod)
		if err != nil {
			return err
		}
		return nil
	}
	parentPath := filepath.Dir(path)
	if !PathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return err
		}
	}
	err = os.WriteFile(path, data, fileMod)
	if err != nil {
		return err
	}
	return nil
}

// WriteFileAsJsonBeauty write json file as 0666 and beauty
func WriteFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return WriteFileAsJson(path, v, os.FileMode(0666), coverage, true)
}
