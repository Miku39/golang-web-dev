/**
Take the previous program and change it so that:
func main uses http.Handle instead of http.HandleFunc
Contstraint: Do not change anything outside of func main
**/

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "index")
}

func dog(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "dog")
}

func me(rw http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(rw, "index.gohtml", "Miku")
	if err != nil {
		log.Fatalln(err)
	}
}
