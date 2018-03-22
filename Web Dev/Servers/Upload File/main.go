package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		// open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// info
		fmt.Println("\n File:", f, "\n header:", h, "\nerr:", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// store on the server
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
		<form method="POST" enctype="multipart/form-data">
			<input type="file" name="q" />
			<button type="submit">Go</button>
		</form>
		<br>
	`+s)
}
