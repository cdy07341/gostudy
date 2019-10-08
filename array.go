package main

import (
	"fmt"
)

func main() {
	//b := [2]int{1,2}
	//fmt.Println(b)
	//指向数组的指针
	//a := [...]int{99:1}
	//var p *[100]int = &a
	//fmt.Println(p)

	//指针数组
	//x,y:=1,2
	//arr := [2]*int{&x, &y}
	//fmt.Println(arr)

	//排序
	//a := [...]int{1,3,5,4,9,6}
	//fmt.Println(a)
	//len := len(a)
	//for i := 0; i < len; i++ {
	//	for j := i + 1; j<len;j++  {
	//		if a[i] < a[j] {
	//			//tmp := a[i]
	//			//a[i] = a[j]
	//			//a[j] = tmp
	//
	//			a[i], a[j] = a[j], a[i]
	//		}
	//	}
	//}
	//
	//fmt.Println(a)

	const (
		a = "a"
		d = "cv"
		b = iota
		c
	)

	const (
		m = "a"
		n = iota
	)

	//fmt.Println(a,b,c,d)
	fmt.Println(m, n)

}
