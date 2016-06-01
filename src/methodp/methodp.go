package methodp

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	name string
}

type Z int

func MethodBasic() {
	var z Z
	z.IncreasePlus(100)
	fmt.Println(z)
}

func (a A) duang() {
	fmt.Println(a.Name, ",Duang!")
}

func (b *B) dang() {
	b.Name = "Chen"
	fmt.Println(b.Name, ",Dang!")
}

func (z *Z) PrintInt() {
	fmt.Println("Print Integer!")
}

func (c *C) duang() {
	c.name = "Vzr"
	fmt.Println(c)
}

func (z *Z) Increase() {
	*z += 100
}

// 虽然Z的底层类型为int，但int并不能直接与Z类型别名进行直接运算
func (z *Z) IncreasePlus(num int) {
	*z += Z(num)
}
