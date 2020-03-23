// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func generate() chan int{
	ch := make(chan int)
	go func() {
		for i:=2; i< 100; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}

func sieve() chan int{
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <- ch
			ch = filter(ch, prime)
			out <- prime
		}

	}()

	return out
}

func filter(in chan int, prime int) chan int{
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				fmt.Println(i)
				//out <- i
			}
		}
	}()

	return out
}
