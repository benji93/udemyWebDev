package main

import (
	"log"
	"os"
	"text/template"
)

type Age struct {
	Age int
}

type Person struct {
	FName string
	LName string
	Age
	Pet
}

type Pet struct {
	Animal string
	FName string
	Age
}

func (p Person) GetAge() int {
	return p.Age.Age
}

func (p Person) GetAgeSecs(x int) int {
	return x * 365 * 24 * 60 * 60
}

func (p Person) FullName() string {
	return p.FName + " " + p.LName
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := Person{
		FName: "Ben",
		LName: "Fell",
		Age: Age{
			24,
		},
		Pet: Pet{
			Animal: "Tarantula",
			FName: "Chipsticks",
			Age: Age{
				5,
			},
		},
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}