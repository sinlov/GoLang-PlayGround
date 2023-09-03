package go_base

import (
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"testing"
)

// 函数是Go里面的核心设计，它通过关键字`func`来声明，它的格式如下：
//
// func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
// 这里是处理逻辑代码
// 返回多个值
// return value1, value2
// }
//
// 上面的代码我们看出
//
// - 关键字`func`用来声明一个函数`funcName`
// - 函数可以有一个或者多个参数，每个参数后面带有类型，通过`,`分隔
// - 函数可以返回多个值
// - 上面返回值声明了两个变量`output1`和`output2`，如果你不想声明也可以，直接就两个类型
// - 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
// - 如果没有返回值，那么就直接省略最后的返回信息
// - 如果有返回值， 那么必须在函数的外层添加return语句

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
	l4g.Debug(i)
	fmt.Println(i)
	i++
	if i < 100 {
		goto Here
	}
	// 用goto跳转到必须在当前函数内定义的标签, 注意大小写
}

func Test_for(t *testing.T) {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	//平行赋值i, j = i+1, j-1 忽略expression1和expression3
	lSum := 1
	for lSum < 1000 {
		lSum += lSum
	}
	fmt.Println("lSum to ", lSum)

	// 也可以完全省略,几乎等于while的用法,当然go是没有while的

	wSum := 1
	for wSum < 1000 {
		wSum += wSum
	}
	fmt.Println("wSum to ", wSum)

	//break操作是跳出当前循环，continue是跳过本次循环
	//当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置
	for index := 10; index > 0; index-- {
		if index == 5 {
			break // 或者continue
		}
		fmt.Println(index)
	}

	// break和continue还可以跟着标签，用来跳到多重循环中的外层循环

	//for配合range可以用于读取slice和map的数据
	rMap := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	for fk, fv := range rMap {
		fmt.Println("map's key:", fk)
		fmt.Println("map's val:", fv)
	}
	//支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值
	eMap := []int{1, 2, 3, 4, 5, 6}
	for _, eV := range eMap {
		fmt.Println("error map val is ", eV)
	}
}

func Test_switch(t *testing.T) {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
	//Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
	//而如果switch没有表达式，它会匹配true

	//Go里面switch默认相当于每个case最后带有break
	// 匹配成功后不会自动向下执行其他case，而是跳出整个switch
	// 但是可以使用fallthrough强制执行后面的case代码

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func sumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func Test_mutli_function(t *testing.T) {
	r1, r2 := sumAndProduct(1, 2)
	fmt.Println(r1, r2)
	//不导出的函数 返回的时候可以不用带上变量名，因为直接在函数里面初始化
	//但如果你的函数是导出的(首字母大写)，官方建议：最好命名返回值

	// Go函数支持变参。接受变参的函数是有着不定数量的参数的 这点更很多别的语言一样就是数组入参
}

//传指针使得多个函数能操作同一个对象

//传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。
//如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。
//go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针

// 简单的一个函数，实现了参数+1的操作
func add1(a *int) int {
	// 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}

func Test_point_add1(t *testing.T) {
	x := 3

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add1(&x) // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"
}
