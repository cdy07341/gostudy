// reflection
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// u := User{1, "OK", 2}
	// info(u)

	//获取匿名字段信息
	// m := Manager{User: User{1, "daoyan", 123}, title: "title value"}
	// t := reflect.TypeOf(m)
	// // fmt.Printf("%#v\n", t.Field(0))
	// fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
	//修改字段值
	// u := User{1, "daoyan", 21}
	// Set(&u)
	// fmt.Println(u)
	//调用方法
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println(name, "my name is ", u.Name)
}

// type Manager struct {
// 	User
// 	title string
// }

// func (u User) Hello() {
// 	fmt.Println("hello world!")
// }

// func info(o interface{}) {
// 	t := reflect.TypeOf(o)
// 	fmt.Println(t.Name())

// 	v := reflect.ValueOf(o)

// 	for i := 0; i < t.NumField(); i++ {
// 		f := t.Field(i)
// 		val := v.Field(i).Interface()
// 		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
// 	}

// 	for i := 0; i < t.NumMethod(); i++ {
// 		m := t.Method(i)
// 		fmt.Println(m.Name, ":", m.Type)
// 	}

// }

// func Set(o interface{}) {
// 	v := reflect.ValueOf(o)
// 	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
// 		fmt.Println("xxx")
// 		return
// 	} else {
// 		v = v.Elem()
// 	}

// 	f := v.FieldByName("Name")
// 	if !f.IsValid() {
// 		fmt.Println("BAD")
// 		return
// 	}

// 	if f.Kind() == reflect.String {
// 		f.SetString("BYEBYE")
// 	}
// }
