package folder_test

import (
	"github.com/sinlov/GoLang-PlayGround/util/folder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentExecPath(t *testing.T) {
	// mock GetCurrentExecPath
	t.Logf("~> mock GetCurrentExecPath")
	// do GetCurrentExecPath
	t.Logf("~> do GetCurrentExecPath")
	// verify GetCurrentExecPath
	execPath, err := folder.GetCurrentExecPath()
	if err != nil {
		t.Error(execPath)
	}
	t.Logf("do TestGetCurrentExecPath test %s , path: %s", t.Name(), execPath)
	assert.NotEqual(t, "", execPath)
}

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

	innerLev1JsonCnt := 5
	innerLev1txtCnt := 2

	innerLev1Folder := folder.PathJoin(testDataPath, "inner_1")
	err = addTextFileByTry(innerLev1Folder, "data", "json", innerLev1JsonCnt)
	if err != nil {
		t.Error(err)
	}
	err = addTextFileByTry(innerLev1Folder, "example", "txt", innerLev1txtCnt)
	if err != nil {
		t.Error(err)
	}

	innerLev2JsonCnt := 4
	innerLev2Folder := folder.PathJoin(innerLev1Folder, "inner_2")
	err = addTextFileByTry(innerLev2Folder, "data", "json", innerLev2JsonCnt)
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

	fileBySuffixJson, err := folder.WalkAllFileBySuffix(testDataPath, "json")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, fileBySuffixJson)
	assert.Equal(t,
		rootLevCnt+innerLev1JsonCnt+innerLev2JsonCnt,
		len(fileBySuffixJson))

	matchJsonPath, err := folder.WalkAllByMatchPath(testDataPath, `.*.json$`, true)

	assert.NotEmpty(t, matchJsonPath)
	assert.Equal(t,
		rootLevCnt+innerLev1JsonCnt+innerLev2JsonCnt,
		len(matchJsonPath))

	matchInnerPath, err := folder.WalkAllByMatchPath(testDataPath, `^.*inner_[0-9]$`, false)
	assert.NotEmpty(t, matchInnerPath)
	assert.Equal(t,
		2,
		len(matchInnerPath))
}
