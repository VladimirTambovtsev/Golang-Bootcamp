package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/products", products)
	http.HandleFunc("/contacts", contacts)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at /index", req.Method, "\n\n")

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func products(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("ma-cookie")
	if err != nil {
		http.SetCookie(res, &http.Cookie{
			Name:  "ma-cookie",
			Value: "Value of this cookie, don't touch it 😱",
		})
		fmt.Fprintln(res, "Your cookie is written, check yo browser and refresh this page: dev tools / application / cookies")
	} else {
		tpl.ExecuteTemplate(res, "products.gohtml", nil)
		fmt.Print(res, "\n\n Your cookie, Sir: ", c)
	}
}

func contacts(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "contacts.gohtml", nil)
}
