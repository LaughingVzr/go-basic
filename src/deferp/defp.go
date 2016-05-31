package deferp

import "fmt"

func DeferBasic() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func DeferFunc() {
	for i := 0; i < 5; i++ {
		fmt.Println("i=", i)
		// 这里由于闭包的出现，方法内的i变量是对外层函数i的一个地址引用，所以defer是在循环结束后才执行的，所以这里的i一直为5
		defer func() {
			fmt.Println(i)
		}()
	}
}

func DeferRecover() {
	printA()
	panicB()
	deferC()
}

func printA() {
	fmt.Println("Func A")
}

func panicB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func deferC() {
	fmt.Println("Defer C")
}

func Complex() {
	// 函数数组
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		// 遵循值copy的原则，这里的i变成了参数
		defer fmt.Println("defer i=", i)
		// 这里使用的defer与闭包的结合使用，使用的i也是外层的地址引用，并且defer是当循环结束后运行，则都为4
		defer func() {
			fmt.Println("defer_closure i=", i)
		}()
		// 这里的i只是一个引用，所以闭包函数内访问的都是i的引用，最终循环完以后调用每一个函数则每个都为4
		fs[i] = func() {
			fmt.Println("closure i=", i)
		}
	}

	// 循环调用闭包函数
	for _, f := range fs {
		f()
	}
}
