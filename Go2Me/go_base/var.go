package go_base

const v1 = 1
const v2 = 2
const v3 = 3

var vname1, vname2, vname3 = v1, v2, v3

// only use in function
//vnamef1, vnamef2, vnamef3 := v1, v2, v3

// 在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量

// _（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃
var _, bitCode = 34, 35

// Go对于已声明但未使用的变量会在编译阶段报错，比如下面的代码就会产生一个错误
func main() {
	//var i int
}

var isActivityVar bool
var enable, disable = true, false

func go_var_test_bool() {
	var available bool
	valid := false
	available = false
	print(available, valid)
}
