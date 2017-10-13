package main

import (
	"log"
	"os"
	"strings"
	"strconv"
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

//  Create a func map to register the function
var fm = template.FuncMap {
	"uc": strings.ToUpper,
	"ft": firstThree,
	"mph": addMph,
}

func init() {
	// New becomes the pointer name (not needed here), then takes the functions in Funcs
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return s
}

func addMph(s int) string {
	t := strconv.Itoa(s)
	t = t + " MPH"
	return t
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

	people := []person{b,s}
	cars := []car{f,h}

	data := struct{
		People []person
		Cars []car
	}{
		people,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}