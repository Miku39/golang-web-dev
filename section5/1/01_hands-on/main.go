/**
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
**/

package main

import (
	"io"
	"net/http"
)

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
	io.WriteString(rw, "hello miku")
}
