package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct { // should be removed to /model
	Email    string
	Password []byte
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/contacts", contacts)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/signin", signin)

	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func contacts(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "contacts.html", nil)
}

func signup(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

func signin(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			log.Println("Alarm! Error logging in")
		}
		email := req.Form.Get("email")
		password := req.Form.Get("password")
		if email == "test@gmail.com" && password == "password" {
			http.Redirect(res, req, "/contacts", http.StatusTemporaryRedirect)
			return
		} else {
			fmt.Println("email is not eq test@gmail.com or password isnt eq to `password`")
		}
	}
	tpl.ExecuteTemplate(res, "signin.html", nil)
}
