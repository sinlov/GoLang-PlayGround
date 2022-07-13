package folder

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestPathFolderList(t *testing.T) {
	// mock PathFolderList
	t.Logf("~> mock PathFolderList")
	// do PathFolderList
	t.Logf("~> do PathFolderList")
	folderList := PathFolderList("../")
	t.Logf("folderList %v", folderList)
	// verify PathFolderList
	assert.NotEqual(t, 0, len(folderList))
}

func TestReadFileAsJson(t *testing.T) {
	// mock ReadFileAsJson

	t.Logf("~> mock ReadFileAsJson")
	// do ReadFileAsJson
	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}
	err := ReadFileAsJson("foo.json", &foo)
	t.Logf("~> do ReadFileAsJson")
	if err != nil {
		t.Error(err)
	}
	// verify ReadFileAsJson
	assert.Equal(t, "name", foo.Bar)
}

func TestWriteFileAsJsonBeauty(t *testing.T) {
	// mock WriteFileAsJsonBeauty

	t.Logf("~> mock WriteFileAsJsonBeauty")
	filePath, err := GetCurrentFilePath()
	if err != nil {
		t.Error(err)
	}
	dir := filepath.Dir(filePath)
	writePath := PathJoin(dir, "bar.json")
	t.Logf("writePath %v", writePath)
	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}
	// do WriteFileAsJsonBeauty
	t.Logf("~> do WriteFileAsJsonBeauty")
	err = WriteFileAsJsonBeauty(writePath, foo, true)
	if err != nil {
		t.Error(err)
	}
	// verify WriteFileAsJsonBeauty
	assert.Nil(t, err)
}
