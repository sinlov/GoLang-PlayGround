package go_base

import (
	"fmt"
	"testing"
)

func Test_if(t *testing.T) {
	if x > 10 {
		fmt.Println("number is max 10")
	} else {
		fmt.Println("number is min 10")
	}

	if y == 3 {
		fmt.Println("The integer is equal to 3")
	} else if y < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}

}

func Test_goto(t *testing.T) {
	i := 1
	Here:
	fmt.Println(i)
	i++
	if i < 100 {
		goto Here
	}
	// 用goto跳转到必须在当前函数内定义的标签, 注意大小写
}

func Test_for(t *testing.T) {
	sum := 0;
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

}
