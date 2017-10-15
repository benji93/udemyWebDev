package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Page struct {
	Title string
	Heading string
	Input string
}

func main() {
	home := Page{
		Title: "Home Page",
		Heading: "New home page",
		Input: "<script>alert()</script>", // This will be escaped
	}

	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}