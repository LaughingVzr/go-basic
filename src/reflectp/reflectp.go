package reflectp

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func (u User) Hello(mgr Manager) {
	fmt.Println(mgr.title)
	fmt.Println(mgr.Id)
	fmt.Println(mgr.Name)
	fmt.Println(mgr.Age)

}

func (u User) HiMulti(i int, str string, fl float64, its []int) {
	fmt.Println(i)
	fmt.Println(str)
	fmt.Println(fl)
	fmt.Println(its)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("非结构类型")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}

func AnymonusColPick() {
	m := Manager{User: User{2, "Vzr", 26}, title: "mananger"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

}

func Invoke() {
	u := User{1, "Laughing", 25}
	InvokeFunc(u)
}

func UserUpdateRef() {
	u := User{1, "Laughing", 25}
	Set(&u)
	fmt.Println(u)
}

/**
* 反射赋值
 */
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if k := v.Kind(); k == reflect.Struct && !v.Elem().CanSet() {
		fmt.Println("不可修改")
		return
	} else {
		v = v.Elem()
	}
	// 取出该字段
	col := v.FieldByName("Name")

	if col.IsValid() && col.Kind() == reflect.String {
		col.SetString("Vzr")
	}
}

func InvokeFunc(o interface{}) {
	v := reflect.ValueOf(o)
	method := v.MethodByName("Hello")
	mgr := Manager{User: User{1, "Laughing", 26}, title: "manager"}
	args := []reflect.Value{reflect.ValueOf(mgr)}
	method.Call(args)

	mthi := v.MethodByName("HiMulti")
	argc := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf("Laughing"), reflect.ValueOf(0.66), reflect.ValueOf([]int{1, 2, 3})}
	mthi.Call(argc)
}
