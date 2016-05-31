package function

import (
	"fmt"
)

func AFunction() {

}

func privateFunction() {

}

func paramFunction(a int, b string, c float64) {

}

func ParamReturnFunc(a int) int {
	a++
	fmt.Println(a)
	return a
}

func MultiReturnFunc(a int) (int, string) {
	a++
	s := string(a)
	return a, s
}

func MultiReturnNameFunc(d int) (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func Closures(str string) func(string) string {
	return func(name string) string {
		return name + "," + str
	}
}

func AnonymousFunc(f func(string)) {
	f("Laughing")
}
