package security_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	defTimeoutSecond = 10
)

var (
	envDebug    = false
	envUserName = ""
	envPassword = ""

	envPluginWebhook = ""
)

func init() {
	envDebug = os.Getenv("ENV_DEBUG") == "true"

	envUserName = os.Getenv("ENV_PLUGIN_USERNAME")
	envPassword = os.Getenv("ENV_PLUGIN_PASSWORD")

	envPluginWebhook = os.Getenv("PLUGIN_WEBHOOK")
}

func envCheck(t *testing.T) bool {
	mustSetEnvList := []string{
		"PLUGIN_WEBHOOK",
	}
	for _, item := range mustSetEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %v, than run test", mustSetEnvList)
			return true
		}
	}

	return false
}

// test case file tools start

var currentFolderAbsPath = ""

func getOrCreateTestDataFolderFullPath() (string, error) {
	if currentFolderAbsPath != "" {
		return currentFolderAbsPath, nil
	}
	currentFolderPath, err := getCurrentFolderPath()
	if err != nil {
		return "", err
	}
	currentFolderAbsPath = path.Join(currentFolderPath, "testdata")
	if !pathExistsFast(currentFolderAbsPath) {
		err := mkdir(currentFolderAbsPath)
		if err != nil {
			currentFolderAbsPath = ""
			return "", err
		}
	}
	return currentFolderAbsPath, nil
}

func goldenDataSaveFast(t *testing.T, data []byte, extraName string) error {
	return goldenDataSave(t, data, extraName, os.FileMode(0766))
}

func goldenDataSave(t *testing.T, data []byte, extraName string, fileMod fs.FileMode) error {
	testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return fmt.Errorf("try goldenDataSave err: %v", err)
	}
	savePath := filepath.Join(testDataFolderFullPath, fmt.Sprintf("%s-%s.golden", t.Name(), extraName))
	err = writeFileByByte(savePath, data, fileMod, true)
	if err != nil {
		return fmt.Errorf("try goldenDataSave at path: %s err: %v", savePath, err)
	}
	return nil
}

func goldenDataReadAsByte(t *testing.T, extraName string) ([]byte, error) {
	testDataFolderFullPath, err := getOrCreateTestDataFolderFullPath()
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}

	savePath := filepath.Join(testDataFolderFullPath, fmt.Sprintf("%s-%s.golden", t.Name(), extraName))

	fileAsByte, err := readFileAsByte(savePath)
	if err != nil {
		return nil, fmt.Errorf("try goldenDataReadAsByte err: %v", err)
	}
	return fileAsByte, nil
}

// getCurrentFolderPath can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// pathExists path exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// pathExistsFast path exists fast
func pathExistsFast(path string) bool {
	exists, _ := pathExists(path)
	return exists
}

// rmDir remove dir by path
func rmDir(path string, force bool) error {
	if force {
		return os.RemoveAll(path)
	}
	exists, err := pathExists(path)
	if err != nil {
		return err
	}
	if exists {
		return os.RemoveAll(path)
	}
	return fmt.Errorf("remove dir not exist path: %s , use force can cover this err", path)
}

// mkdir will use FileMode 0766
func mkdir(path string) error {
	err := os.MkdirAll(path, os.FileMode(0766))
	if err != nil {
		return fmt.Errorf("fail MkdirAll at path: %s , err: %v", path, err)
	}
	return nil
}

// readFileAsByte read file as byte array
func readFileAsByte(path string) ([]byte, error) {
	exists, err := pathExists(path)
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

// readFileAsJson read file as json
func readFileAsJson(path string, v interface{}) error {
	fileAsByte, err := readFileAsByte(path)
	err = json.Unmarshal(fileAsByte, v)
	if err != nil {
		return fmt.Errorf("path: %s , read file as json err: %v", path, err)
	}
	return nil
}

// writeFileByByte write bytes to file
// path most use Abs Path
// data []byte
// fileMod os.FileMode(0666) or os.FileMode(0644)
// coverage true will coverage old
func writeFileByByte(path string, data []byte, fileMod fs.FileMode, coverage bool) error {
	if !coverage {
		exists, err := pathExists(path)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("not coverage, which path exist %v", path)
		}
	}
	parentPath := filepath.Dir(path)
	if !pathExistsFast(parentPath) {
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

// writeFileAsJson write json file
// path most use Abs Path
// v data
// fileMod os.FileMode(0666) or os.FileMode(0644)
// coverage true will coverage old
// beauty will format json when write
func writeFileAsJson(path string, v interface{}, fileMod fs.FileMode, coverage, beauty bool) error {
	if !coverage {
		exists, err := pathExists(path)
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
		return writeFileByByte(path, str.Bytes(), fileMod, coverage)
	}
	return writeFileByByte(path, data, fileMod, coverage)
}

// writeFileAsJsonBeauty write json file as 0766 and beauty
func writeFileAsJsonBeauty(path string, v interface{}, coverage bool) error {
	return writeFileAsJson(path, v, os.FileMode(0766), coverage, true)
}

// test case file tools end
