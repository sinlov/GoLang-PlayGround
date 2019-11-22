package go_runtime

import (
	"crypto/sha1"
	"fmt"
	"strconv"
)

type Obj struct {
	Data []byte
}

func getObjHashCode(obj *Obj) (int, error) {
	fmt.Printf("%p\n", obj)
	hash := sha1.New()
	sum := hash.Sum(obj.Data)
	sumStr := fmt.Sprintf("%x", sum)
	code, err := strconv.Atoi(sumStr)
	if err != nil {
		return -1, err

	}
	return code, nil
}
