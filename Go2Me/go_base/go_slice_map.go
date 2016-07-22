package go_base

import "fmt"

// base array
var arr [10]int  // 声明了一个int类型的数组

func test_array() {
	arr[0] = 0
	arr[3] = 3
	fmt.Printf("The first element is %d\n", arr[0])
	//返回未赋值的最后一个元素，默认返回0
	fmt.Printf("The last element is %d\n", arr[9])
	// 快速声明
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Println(a, b, c)
	//多维数组
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	// 声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doubleArray, easyArray)
}

//由于长度也是数组类型的一部分，因此 [3]int 与 [4]int 是不同的类型，数组也就不能改变长度。
//数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。

//这个时候需要动态数组 slice

//slice并不是真正意义上的动态数组，而是一个引用类型。
//slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

var mSlice []int

// slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i
var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
var a, b []byte

func test_slice() {
	sSlice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(sSlice)
	a = ar[2:5]
	b = ar[3:5]
	//slice有一些简便的操作
	ar[:3] == ar[0:3]
	arlen := len(ar)
	ar[9:] == ar[9:arlen]
	ar[:] == ar[0:arlen]

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
	fmt.Println(bSlice)

	//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值
	//例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变

	//内置函数
	len(ar) //获取长度
	cap(ar) //获取最大容量
	append(ar) //追加一个或者多个元素，然后返回一个和slice一样类型的slice
	var br []byte
	copy(br, ar) //从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

	// 三个参数的slice go 1.2 开始支持

	tSlice := array[2:4:7]
	fmt.Println(tSlice)
	// 产生的新的slice就没办法访问最后的三个元素
}

// append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。
// 但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
// 返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响


func test_map()  {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	// 另一种map的声明方式
	numbers = make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3

	// 用map过程中需要注意的几点

	//map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
	//map的长度是不固定的，也就是和slice一样，也是一种引用类型
	//内置的len函数同样适用于map，返回map拥有的key的数量
	//map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
	//map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制

	//map的初始化可以通过key:val的方式初始化值

	// 初始化一个字典
	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	cSharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", cSharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")  // 删除key为C的元素

	//如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了
}
