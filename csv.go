package main

import (
	"encoding/csv"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
)

func main() {
	fileName := "a.csv"
	file, e := os.Open(fileName)
	if nil != e {
		log.Fatal(e.Error())
	}

	reader := csv.NewReader(file)

	line, err := reader.ReadAll()
	if nil != err {
		log.Fatal(err.Error())
	}

	fmt.Println(line)

}
