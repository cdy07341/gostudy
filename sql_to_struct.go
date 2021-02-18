package main

import (
	"flag"
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	table := flag.String("table", "", "table name")
	flag.Parse()
	if "" != *table {
		err := converter.NewTable2Struct().
			SavePath("./" + *table + ".go").
			Dsn("dev:1233456@tcp(10.227.20.144:8306)/web?charset=utf8").
			Table(*table).
			EnableJsonTag(true).
			TagKey("ddb").
			Run()
		fmt.Println(err)
	} else {
		fmt.Println("table empty")
	}

}