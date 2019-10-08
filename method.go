// method
package main

import (
	"fmt"
)

func main() {
	//method value
	// a := A{}
	// a.Print()

	// b := B{}
	// b.Print()

	//method expression
	// (*A).Print(&a)

	//int 增加print方法
	// var a TZ
	// a.Print()

	//作业
	var a TZ
	a.inCrease(100)
	fmt.Println(a)
}

// type A struct {
// 	Name string
// }

// type B struct {
// 	Name string
// }

// func (a A) Print() {
// 	fmt.Println("A")
// }

// func (b B) Print() {
// 	fmt.Println("B")
// }

// type TZ int

// func (tz *TZ) Print() {
// 	fmt.Println("TZ")
// }

type TZ int

func (a *TZ) inCrease(num int) {
	for i := 1; i < 5; i++ {
		*a += TZ(1)
		// fmt.Println(a)
	}
}
