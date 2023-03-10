/**
ListenAndServe on port 8080 of localhost
For the default route "/" Have a func called "foo" which writes to the response "foo ran"
For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response
"This is from dog" and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
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
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<p>foo ran</p>`)
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
