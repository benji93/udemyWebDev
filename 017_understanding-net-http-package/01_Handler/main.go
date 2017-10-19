package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Do you want any code in this func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}