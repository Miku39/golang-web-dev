// Using cookies, track how many times a user has been to your website domain.

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("count")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "count",
			Value: "0",
			Path:  "/",
		})

	}

	numInt, _ := strconv.Atoi(cookie.Value)
	numInt++
	cookie.Value = strconv.Itoa(numInt)

	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Visit Count:", cookie.Value)
}
