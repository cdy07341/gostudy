package main

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/util/log"
	"io"
	"os"
)

func main() {
	written, err := fileCopy("testBak.log", "test.log")
	if nil != err {
		log.Fatal(err.Error())
	}
	fmt.Println(written)
}

func fileCopy(dstName, srcName string) (written int64, err error) {
	srcFile, err := os.Open(srcName)
	if nil != err {
		log.Fatal(err.Error())
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err {
		log.Fatal(err.Error())
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
