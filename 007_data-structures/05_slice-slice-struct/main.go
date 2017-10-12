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

type car struct {
	Name string
	Speed int
}

type data struct {
	People []person
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	b := person{
		"Ben", 
		"Fell",
	}

	s := person{
		"Sarah",
		"Laybourn",
	}

	f := car{
		"Ferrari",
		180,
	}

	h := car{
		"Honda",
		90,
	}

	people := []person{b, s}
	transport := []car{f, h}

	data := data{people,transport}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}