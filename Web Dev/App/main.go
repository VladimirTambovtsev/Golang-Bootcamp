package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var tpl *template.Template // postgres
var db *sql.DB             // db

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct { // should be removed to /model
	Email    string
	Password []byte
}

func SetDatabase(database *sql.DB) {
	db = database
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://vladimir:351444@localhost/consumer_complaints?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to DB: %v", err))
	}
	SetDatabase(db)
	return db
}

func main() {
	db := connectDb()
	defer db.Close()
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
	res.Header().Add("Content-Type", "text/html") // doesn't important
	tpl.ExecuteTemplate(res, "signin.html", nil)
}
