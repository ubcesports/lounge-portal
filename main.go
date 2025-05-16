package main

import (
	"fmt"
	"net/http"
	"github.com/a-h/templ"
)

func main() {
	component := hello("Hello World")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
