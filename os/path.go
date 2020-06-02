package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//fmt.Println(string(os.PathSeparator))
	//paths := "/home/polaris/studygolang/"
	//fmt.Println(path.Dir(paths))
	//fmt.Println(path.Base(paths))
	//fmt.Println(filepath.Dir(paths))

	//fmt.Println(filepath.Clean("/../tmp.txt"))

	filepath.Walk("../net/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fmt.Println("file:", info.Name(), "in directory:", path)

		return nil
	})
}
