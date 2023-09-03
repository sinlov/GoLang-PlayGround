package go_base

import (
	"fmt"
	"testing"
)

// 写Go代码的时候经常用到import这个命令用来导入包文件

func Test_fmt(t *testing.T) {
	// 其实是去GOROOT环境变量指定目录下去加载该模块
	fmt.Println("this is fmt")
}

// 另外的导入方式
// 相对路径 import “./model” //当前文件同一目录的model目录 ,不推荐
// 绝对路径 import “shorturl/model”

// 导入时的特别操作

// 点操作
// . "fmt"
// 点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名
// 也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")

// 别名操作
// f "fmt"
// 别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")

// _操作
// _ "github.com/ziutek/mymysql/godrv"
// _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
