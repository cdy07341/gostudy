package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	//close(ch)
	select {
	case v := <-ch:
		fmt.Println("dfdfd", v)
	default:
		fmt.Println("default")
	}
}
