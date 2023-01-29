package folder_test

import (
	"fmt"
	"github.com/sinlov/GoLang-PlayGround/util/folder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFileAsJson(t *testing.T) {
	// mock ReadFileAsJson

	t.Logf("~> mock ReadFileAsJson")
	// do ReadFileAsJson
	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}
	err := folder.ReadFileAsJson("foo.json", &foo)
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
	filePath, err := folder.GetCurrentFilePath()
	if err != nil {
		t.Error(err)
	}
	dir := folder.PathParent(filePath)
	writePath := folder.PathJoin(dir, "bar.json")

	t.Logf("writePath %v", writePath)
	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}
	foo.Foo = 2
	foo.Bar = "bar"
	// do WriteFileAsJsonBeauty
	t.Logf("~> do WriteFileAsJsonBeauty")

	err = folder.WriteFileAsJsonBeauty(writePath, foo, true)
	if err != nil {
		t.Error("not cover write path not exist")
	}

	writePath = folder.PathJoin(dir, "testdata", "bar.json")

	err = folder.Mkdir(folder.PathParent(writePath))
	if err != nil {
		t.Error(err)
	}
	err = folder.WriteFileAsJsonBeauty(writePath, foo, true)
	if err != nil {
		t.Error(err)
	}
	// verify WriteFileAsJsonBeauty
	assert.Nil(t, err)
	err = folder.WriteFileAsJsonBeauty(writePath, foo, false)
	if err == nil {
		t.Error("not cover write path not exist")
	}
}

func TestPathParent(t *testing.T) {
	// mock PathParent
	pathOne := "foo/bar/foo"
	t.Logf("~> mock PathParent")
	// do PathParent
	t.Logf("~> do PathParent")
	// verify PathParent
	assert.Equal(t, "foo/bar", folder.PathParent(pathOne))

	assert.Equal(t, "/abc/def", folder.PathParent("/abc/def/ghf"))
	assert.Equal(t, "../abc/def", folder.PathParent("../abc/def/ghf"))
}

func addTextFileByTry(targetDir, fileHead, suffix string, cnt int) error {

	if !folder.PathExistsFast(targetDir) {
		err := folder.Mkdir(targetDir)
		if err != nil {
			return err
		}
	}

	var foo struct {
		Foo int    `json:"foo"`
		Bar string `json:"bar"`
	}

	for i := 0; i < cnt; i++ {
		fName := fmt.Sprintf("%s_%d.%s", fileHead, i, suffix)
		newJsonPath := folder.PathJoin(targetDir, fName)
		foo.Foo = i
		errJsonWrite := folder.WriteFileAsJsonBeauty(newJsonPath, foo, true)
		if errJsonWrite != nil {
			return errJsonWrite
		}
	}
	return nil
}

func TestFolder(t *testing.T) {
	currentFilePath, err := folder.GetCurrentFilePath()
	if err != nil {
		t.Error(err)
	}
	currentFolder := folder.PathParent(currentFilePath)
	testDataPath := folder.PathJoin(currentFolder, "testdata")

	err = folder.RmDirForce(testDataPath)
	if err != nil {
		t.Error(err)
	}
	err = folder.Mkdir(testDataPath)
	if err != nil {
		t.Error(err)
	}

	rootLevCnt := 3

	err = addTextFileByTry(testDataPath, "data", "json", rootLevCnt)
	if err != nil {
		t.Error(err)
	}

	innerLev1Folder := folder.PathJoin(testDataPath, "inner_1")
	err = addTextFileByTry(innerLev1Folder, "data", "json", 5)
	if err != nil {
		t.Error(err)
	}
	err = addTextFileByTry(innerLev1Folder, "example", "txt", 5)
	if err != nil {
		t.Error(err)
	}

	innerLev2Folder := folder.PathJoin(innerLev1Folder, "inner_2")
	err = addTextFileByTry(innerLev2Folder, "data", "json", 5)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, folder.PathIsDir(innerLev1Folder))
	assert.False(t, folder.PathIsFile(innerLev1Folder))

	rootFileList := folder.PathFileList(testDataPath)
	assert.NotEmpty(t, rootFileList)
	assert.Equal(t, rootLevCnt, len(rootFileList))
	rootFirstFilePath := folder.PathJoin(testDataPath, rootFileList[0])
	t.Logf("first root path: %s", rootFirstFilePath)
	assert.True(t, folder.PathIsFile(rootFirstFilePath))
	assert.False(t, folder.PathIsDir(rootFirstFilePath))

	rootFolderList := folder.PathFolderList(testDataPath)
	assert.NotEmpty(t, rootFolderList)
	assert.Equal(t, 1, len(rootFolderList))
}
