package folder_test

import (
	"fmt"
	"github.com/sinlov/GoLang-PlayGround/util/folder"
)

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
