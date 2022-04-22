package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/html", indexHtml)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "simple web")
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("xxxx")
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}
