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
		"Leader": "Ben",
		"Actual Leader": "Sarah",
		"Co-Leader": "Katie",
		"WeaponMaster": "Tom",
	}

	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		log.Fatalln(err)
	}
}