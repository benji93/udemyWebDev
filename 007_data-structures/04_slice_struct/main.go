package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	FirstName string
	LastName string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ben := person{
		"Ben",
		"Fell",
	}

	sarah := person{
		"Sarah",
		"Laybourn",
	}

	tom := person{
		"Tom",
		"Thorp",
	}

	people := []person{ben, sarah, tom}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		log.Fatalln(err)
	}
}