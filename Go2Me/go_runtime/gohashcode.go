package go_runtime

import (
	"fmt"
	"crypto/sha1"
	"strconv"
)

type Obj struct {


}

func getObjHashCode(obj *Obj) (int, error) {
	fmt.Printf("%p\n", obj)
	hash := sha1.New()
	sum := hash.Sum([]byte(&obj))
	sumStr := fmt.Sprintf("%x", sum)
	code, err := strconv.Atoi(sumStr)
	if err != nil {
		return -1, err

	}
	return code, nil
}
