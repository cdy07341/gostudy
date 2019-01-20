// struct
package main

import (
	"fmt"
)

func main() {

	// person := Person{
	// 	Name: "fff",
	// 	Age:  444,
	// }
	// person.Age = 18
	// person.Name = "dddd"
	// fmt.Println(person)
	// A(&person)
	// fmt.Println(person)

	//匿名结构
	// a := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "fff",
	// 	Age:  444,
	// }

	// fmt.Println(a)

	//组合
	// a := teacher{
	// 	Name:  "cdy",
	// 	Age:   10,
	// 	human: human{Sex: 1},
	// }

	// b := student{
	// 	Name:  "daoyan",
	// 	Age:   11,
	// 	human: human{Sex: 0},
	// }

	// b.human.Sex = 100

	// fmt.Println(a)
	// fmt.Println(b)

	//匿名字段
	// type test struct {
	// 	string
	// 	int
	// }

	// a := test{"name", 111}
	// fmt.Println(a.int)

}

// type Person struct {
// 	Name string
// 	Age  int
// }

// func A(obj *Person) {
// 	obj.Age = 555
// 	fmt.Println(obj)
// }

// type human struct {
// 	Sex int
// }

// type teacher struct {
// 	human
// 	Name string
// 	Age  int
// }

// type student struct {
// 	human
// 	Name string
// 	Age  int
// }
