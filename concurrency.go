// concurrency
package main

import "fmt"

func main() {
	// go Go()
	// time.Sleep(2 * time.Second)

	//c := make(chan bool)
	//go func() {
	//	fmt.Println("GO GO GO!!!")
	//	c <- true
	//}()
	//<-c

	//select
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool, 2)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1")
					o <- true
					break
				}
				fmt.Println("c1:", v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2")
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 2
	c2 <- "hello"

	close(c1)
	//close(c2)

	for i := 0; i < 5; i++ {
		out := <-o
		fmt.Println(out)
	}

}

// func Go() {
// 	fmt.Println("GO GO GO!!!")
// }
