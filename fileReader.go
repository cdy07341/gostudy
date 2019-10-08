package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	title    string
	price    float64
	quantity int
}

func main() {
	products := make([]Product, 1)

	fileName := "a.log"
	file, err := os.Open(fileName)
	if nil != err {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		product := new(Product)
		line, _, err := reader.ReadLine()
		if nil != err {
			break
		}
		arrVal := strings.Split(string(line), ";")
		title := arrVal[0]
		price, _ := strconv.ParseFloat(arrVal[1], 32)
		quantity, _ := strconv.Atoi(arrVal[2])

		product.title = title
		product.price = price
		product.quantity = quantity
		if "" == products[0].title {
			products[0] = *product
		} else {
			products = append(products, *product)
		}
	}

	fmt.Println("We have read the following books from the file: ")
	fmt.Println(products)

}
