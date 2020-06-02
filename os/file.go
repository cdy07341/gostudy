package main

import (
	"fmt"
	"os"
)

func main() {
	file,_:= os.Open("tmp.txt")
	fileInfo, _ := file.Stat()
	fmt.Printf("%+v",fileInfo)
	info, _ := os.Stat("tmp.txt")
	fmt.Println(info.Name())
}
