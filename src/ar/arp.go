package ar

import "fmt"

func ArrNormal() {
	/**这里数组的长度也作为了数组的一部分。
	  所以长度不同但类型相同的数组
	  不能使用=直接赋值*/
	var a [2]int
	var b [2]int
	b = a
	fmt.Println(b)
}

func ArrUndefLength() {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{29: 100}
	fmt.Println(a)
	fmt.Println(b)
}

func ArrPointer() {
	// 指向数组的指针
	a := [...]int{99: 1}
	var p *[100]int = &a
	fmt.Println(p)

	// 指针数组
	x, y := 1, 2
	b := [...]*int{&x, &y}
	fmt.Println(b)
}

func ArrCompare() {
	var a [2]int
	var b [2]int
	fmt.Println(a == b)
}

func ArrNew() {
	p := new([10]int)
	fmt.Println(&p)

	// 直接使用数组来操作索引与值
	a := [10]int{}
	a[1] = 2
	fmt.Println(a)

	// 使用数组指针来操作索引与值
	p[1] = 2
	fmt.Println(p)
}

func BubbleSort() {
	// 中间变量
	var tmp int
	a := [...]int{899, 56, 1027, 45, 23, 439, 603, 211, 532}
	// 输出未排序数组
	fmt.Println("排序前：", a)
	length := len(a)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if a[i] > a[j] {
				tmp = a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}
	// 输出排序后数组
	fmt.Println("排序后：", a)
}
