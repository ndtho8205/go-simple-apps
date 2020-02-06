package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Golang")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}
