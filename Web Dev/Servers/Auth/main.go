package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}
