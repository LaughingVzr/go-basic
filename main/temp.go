package main

import (
	"fmt"
	rp "reflectp"
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func main() {
	rp.Invoke()
}

func saveUnit() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}

func getPointr() {
	a := 1
	a--
	var p *int = &a
	fmt.Println(*p)
}

func partVar() {
	a := 10
	// 这里的a会使用局部的a
	if a := 1; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
}

func switchP() {
	// 在分支语句中定义的变量只是局部变量
	switch a := 1; {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}

func gogogo() {
	//break continue goto都可以结合标签使用
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("跳出了循环")
LABEL2:
	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			continue LABEL2
		}
	}
	fmt.Println("OK")
}

func gotogoto() {
	//标签放在goto后，以防死循环
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL
		}
	}
	fmt.Println("跳出了循环")
}
