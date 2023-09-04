package folder

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// GetCurrentExecPath get exec path
func GetCurrentExecPath() (string, error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return s, nil
	//i := strings.LastIndex(s, "\\")
	//p := string(s[0 : i+1])
	//return p, nil
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

// Mkdir will use FileMode 0766
func Mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// RmDir remove dir by path
func RmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := PathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// RmDirForce remove dir force
func RmDirForce(path string) error {
	return RmDir(path, true)
}

// PathIsDir exists path is dir or false
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// PathIsFile exists path is file or false
func PathIsFile(path string) bool {
	fi, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// PathJoin is filepath.Join()
func PathJoin(elem ...string) string {
	return filepath.Join(elem...)
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

	fileInfos, err := os.ReadDir(path)
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

	fileInfos, err := os.ReadDir(path)
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

// ReadFileAsByte read file as byte array
func ReadFileAsByte(path string) ([]byte, error) {
	exists, err := PathExists(path)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("path not exist %v", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("path: %s read err: %v", path, err)
	}

	return data, nil
}

// ReadFileAsJson read file as json
func ReadFileAsJson(path string, v interface{}) error {
	fileAsByte, err := ReadFileAsByte(path)
	if err != nil {
		return fmt.Errorf("path: %s , read file as byte err: %v", path, err)
	}
	err = json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
}

// WriteFileByByte write bytes to file
// path most use Abs Path
// data []byte
// fileMod os.FileMode(0666) or os.FileMode(0644)
// coverage true will coverage old
func WriteFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := PathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !PathExistsFast(parentPath) {
		err := os.MkdirAll(parentPath, fileMod)
		if err != nil {
			return fmt.Errorf("can not WriteFileByByte at new dir at mode: %v , at parent path: %v", fileMod, parentPath)
		}
	}
	err := os.WriteFile(path, data, fileMod)
	if err != nil {
		return fmt.Errorf("write data at path: %v, err: %v", path, err)
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
		return WriteFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return WriteFileByByte(path, data, fileMod, coverage)
}

// WriteFileAsJsonBeauty write json file as 0666 and beauty
func WriteFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return WriteFileAsJson(path, v, os.FileMode(0666), coverage, true)
}
