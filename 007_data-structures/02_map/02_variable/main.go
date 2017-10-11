package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	people := map[string]string{
		"Best": "Ben",
		"OtherBest": "Sarah",
		"Cool": "Katie",
		"Guns": "Tom",
	}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		log.Fatalln(err)
	}
}