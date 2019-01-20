// index
package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// y := closure(10)
	// fmt.Println(y(1))

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)

	}

}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}

}
