package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/items/", items)
	//http.HandleFunc("/template/", template)
	http.HandleFunc("/about/", about)
	http.HandleFunc("/contact/", contact)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Main Page")
}

func items(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Items Page")
}

func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "About Page")
}

func contact(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Contact Page")
}

// func template(res http.ResponseWriter, req *http.Request) {
// 	tmpl, err := template.ParseFiles("index.html")

// 	if err != nil {
// 		log.Fatalln("error parsing template", err)
// 	}

// 	err := tpl.ExecuteTemplate(res, "index.gohtml", "Vlad")
// 	if err != nil {
// 		log.Fatalln("error executing template", err)
// 	}
// }
