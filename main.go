package main

import "fmt"

type Student struct {
	Name string
}

func (s Student) getName() string {
	return s.Name
}

func main() {
	var s Student

	name := s.getName()

	fmt.Println(name)
}