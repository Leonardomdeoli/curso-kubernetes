package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	w.Write([]byte("<h1>Nome</h1>" + name + "</br><h1>Idade</h1>" + age))
}