package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (this *Page) save() error {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page) load(title string) error {
	var err error
	this.Title = title
	this.Body, err = ioutil.ReadFile(title)

	return err
}

func main() {
	page := &Page{
		Title: "test.log",
		Body:  []byte("# Page\n## Section1\nThis is section1."),
	}

	page.save()

	var newPage Page
	newPage.load("test.log")
	fmt.Println(string(newPage.Body))

}
