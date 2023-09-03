package go_base

// 布尔值的类型为bool，值是true或false，默认为false

var isActivity bool = true

// 忽略类型的声明
var enabled, disabled = true, false

func go_base_type_test() {
	var isActivity bool // 一般声明
	valid := false      // 简短声明
	isActivity = true   // 赋值操作
	print(isActivity)
	print(valid)
}
