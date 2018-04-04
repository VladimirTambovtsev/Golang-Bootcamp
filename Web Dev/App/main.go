package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
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
	tpl.ExecuteTemplate(res, "signin.html", nil)
}
