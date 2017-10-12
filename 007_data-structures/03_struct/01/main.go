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
	p := person{
		FirstName: "Ben",
		LastName: "Fell",
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}