// panicAndRecover
package main

import (
	"fmt"
)

func main() {
	// A()
	// B()
	// C()

	//作业
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer_closuer i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closuer i = ", i)
		}
	}

	for _, f := range fs {
		f()
	}

}

func A() {
	fmt.Println("abb c")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover B")
		}
	}()
	fmt.Println("b")
	panic("bb")
}

func C() {
	fmt.Println("c")
}
