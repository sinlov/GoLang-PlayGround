package go_base

import "fmt"

// go use "" or `` to mark sting

// 声明变量为字符串的一般方法
var frenchHello string
// 声明了一个字符串变量，初始化为空字符串
var emptyString string = ""

//声明一个多行的字符串
//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出
var mutli_line_string = `
	my json
	new line
`

func go_base_test() {
	no, yes, maybe := "no", "yes", "maybe"  // 简短声明，同时声明多个变量
	japaneseHello := "Konichiwa"  // 同上
	frenchHello = "Bonjour"  // 常规赋值
	print(no, yes, maybe, japaneseHello)
}

// Go中字符串是不可变的

func go_base_err_string_change()  {
	var s string = "hello"
	//s[0] = 'c' // code worry cannot assign to s[0]
	print(s)
}

/**
 * if you want to change string, you must change to []byte then cast back
 */
func go_base_string_change() {
	s := "hello"
	c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c)  // 再转换回 string 类型
	fmt.Printf("%s\n", s2)
}

func go_base_string_substring() {
	// use + to plus string
	s := "hello,"
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)

	// 字符串虽不能更改，但可进行切片操作
	forS := "hello"
	forS = "c" + forS[1:]
	fmt.Printf("%s\n", forS)
}

