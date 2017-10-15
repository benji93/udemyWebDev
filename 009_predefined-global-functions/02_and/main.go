package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	FName string
	LName string
	IsAdmin bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := User{"Ben","Fell",false}
	u2 := User{"Sarah","Laybourn",true}
	u3 := User{"","Thorp",true}

	users := []User{u1,u2,u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}