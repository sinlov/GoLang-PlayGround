package go_base

import (
	"fmt"
	"testing"
)

// 声明新的类型，作为其它类型的属性或字段的容器

type person struct {
	name string
	age  int
}

func Test_person(t *testing.T) {
	var P person
	P.name = "myName"
	P.age = 233
	fmt.Println("This persion p name is ", P.name)

	fP := person{"Tom", 25}
	fvP := person{age: 24, name: "Tom"}
	// 这个时候 pP 的类型是 *person
	pP := new(person)
	fmt.Println(fP.name)
	fmt.Println(fvP.name)
	fmt.Println(pP)
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		// 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func Test_older_person(t *testing.T) {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{age: 25, name: "Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	tb_Older, tb_diff := older(tom, bob)
	tp_Older, tp_diff := older(tom, paul)
	bp_Older, bp_diff := older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}
