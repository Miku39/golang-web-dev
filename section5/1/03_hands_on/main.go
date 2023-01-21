/**
Take the previous program in the previous folder and change it so that:
a template is parsed and served
you pass data into the template
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
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

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
