package main

import "fmt"

func main() {
	name := "Ben Fell"

	tpl := `
	<!DOCTYPE html>
	<html>
	<head lang="en">
	<meta charset="UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}