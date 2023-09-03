package go_base

import (
	"fmt"
	"os"
	"testing"
)

//即延迟（defer）语句，你可以在函数中添加多个defer语句。
//当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。
//特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题

func readWrite() bool {
	src, err := os.Open("file")
	defer src.Close()
	if err != nil {
		return false
	}
	//if err.Error() {
	//	return false
	//}
	return true
}

//defer是采用后进先出模式

func Test_defer_mode(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
