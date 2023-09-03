package go_base

//在Go语言中，同时声明多个常量、变量，或者导入多个包时，可采用分组的方式进行声明
import (
	"fmt"
	"os"
)

const (
	iConst      = 100
	piConst     = 3.1415
	prefixConst = "Go_"
)

var (
	iVar      int
	piVar     float32
	prefixVar string
)

func go_base_tips() {
	fmt.Println(os.Args)
}

//关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	eConst, fConst, gConst = iota, iota, iota //e=0,f=0,g=0 iota在同一行值相同
)

const (
	a       = iota //   a = 0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)
