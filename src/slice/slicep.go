package slice

import "fmt"

func CutSlice() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	fmt.Println(string(a))

	sa := a[2:5]
	fmt.Println(string(sa))

	fmt.Println(string(a[3:5]))
}

func MadeSlice() {
	s1 := make([]int, 3, 10)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1)
}

func Reslice() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	fmt.Println(len(sa), cap(sa))
	sb := sa[1:]
	fmt.Println(string(sb))
}

func AppendSlice() {
	a := make([]int, 3, 6)
	fmt.Printf("%v %p\n", a, a)

	a = append(a, 1, 2, 3)
	fmt.Printf("%v %p\n", a, a)

	a = append(a, 4)
	fmt.Printf("%v %p\n", a, a)

	fmt.Println("长度为：", len(a), "，容量为：", cap(a))
}

func CopySlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	copy(s2[1:2], s1[2:3])
	fmt.Println(s2)

	s3 := s1[:]
	fmt.Println(s3)
}
