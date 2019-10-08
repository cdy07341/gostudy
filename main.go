package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var total, wordTotal, lineTotal int

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inputReader.ReadString('\n')
		if nil != err {
			fmt.Println("error")
		}

		if "S\n" == input {
			fmt.Printf("Number of characters: %d\n", total)
			fmt.Printf("Number of words: %d\n", wordTotal)
			fmt.Printf("Number of lines: %d\n", lineTotal)
			os.Exit(0)
		}
		Counter(input)
	}

}

func Counter(input string) {
	total += len(input) - 2
	wordTotal += len(strings.Fields(input))
	lineTotal++
}
